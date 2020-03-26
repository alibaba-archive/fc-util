<?php

namespace AlibabaCloud\Tea\FCUtils;

class Auth
{
    private $access_key_id;
    private $access_key_secret;
    private $security_token;

    public function __construct($ak_id, $ak_secret, $ak_secret_token = '')
    {
        $this->access_key_id     = trim($ak_id);
        $this->access_key_secret = trim($ak_secret);
        $this->security_token    = trim($ak_secret_token);
    }

    public function getSecurityToken()
    {
        return $this->security_token;
    }

    public function signRequest($method, $pathname, $headers, $unescaped_queries = null)
    {
        $unescaped_path = $this->unescape($pathname);
        /*
        Sign the request. See the spec for reference.
        https://help.aliyun.com/document_detail/52877.html
        @param $method: method of the http request.
        @param $headers: headers of the http request.
        @param $unescaped_path: unescaped path without queries of the http request.
        @return: the signature string.
         */

        $content_md5        = isset($headers['content-md5']) ? $headers['content-md5'] : '';
        $content_type       = isset($headers['content-type']) ? $headers['content-type'] : '';
        $date               = isset($headers['date']) ? $headers['date'] : '';
        $canonical_headers  = $this->buildCanonicalHeaders($headers);
        $canonical_resource = $unescaped_path;

        if (!empty($unescaped_queries)) {
            $canonical_resource = $this->getSignResource($unescaped_path, $unescaped_queries);
        }

        $string_to_sign = implode(
            "\n",
            [strtoupper($method), $content_md5, $content_type, $date, $canonical_headers . $canonical_resource]
        );
        $h              = hash_hmac('sha256', $string_to_sign, $this->access_key_secret, true);

        return 'FC ' . $this->access_key_id . ':' . base64_encode($h);
    }

    private function buildCanonicalHeaders($headers)
    {
        /*
        @param $headers: array
        @return: $Canonicalized header string.
        @return: String
         */
        $canonical_headers = [];
        foreach ($headers as $k => $v) {
            $lower_key = trim(strtolower($k));
            if ('x-fc-' === substr($lower_key, 0, 5)) {
                $canonical_headers[$lower_key] = $v;
            }
        }
        ksort($canonical_headers);
        $canonical = '';
        foreach ($canonical_headers as $k => $v) {
            $canonical = $canonical . $k . ':' . $v . "\n";
        }

        return $canonical;
    }

    private function getSignResource($unescaped_path, $unescaped_queries)
    {
        if (!\is_array($unescaped_queries)) {
            throw new \Exception('`array` type required for queries');
        }

        $params = [];
        foreach ($unescaped_queries as $key => $values) {
            if (\is_string($values)) {
                $params[] = sprintf('%s=%s', $key, $values);

                continue;
            }
            if (\count($values) > 0) {
                foreach ($values as $value) {
                    $params[] = sprintf('%s=%s', $key, $value);
                }
            } else {
                $params[] = (string) $key;
            }
        }
        ksort($params);

        return $unescaped_path . "\n" . implode("\n", $params);
    }

    private function unescape($str)
    {
        $ret = '';
        $len = \strlen($str);
        for ($i = 0; $i < $len; ++$i) {
            if ('%' == $str[$i] && 'u' == $str[$i + 1]) {
                $val = hexdec(substr($str, $i + 2, 4));
                if ($val < 0x7f) {
                    $ret .= \chr($val);
                } elseif ($val < 0x800) {
                    $ret .= \chr(0xc0 | ($val >> 6)) .
                            \chr(0x80 | ($val & 0x3f));
                } else {
                    $ret .= \chr(0xe0 | ($val >> 12)) .
                            \chr(0x80 | (($val >> 6) & 0x3f)) .
                            \chr(0x80 | ($val & 0x3f));
                }

                $i += 5;
            } elseif ('%' == $str[$i]) {
                $ret .= urldecode(substr($str, $i, 3));
                $i += 2;
            } else {
                $ret .= $str[$i];
            }
        }

        return $ret;
    }
}
