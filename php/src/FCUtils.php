<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\Tea\FCUtils;

use AlibabaCloud\Tea\Request;

class FCUtils
{
    /**
     * @param array $body
     *
     * @throws \Exception
     *
     * @return string
     */
    public static function getContentMD5($body)
    {
        $string = implode('', array_map('chr', $body));

        return md5($string);
    }

    /**
     * @param array $body
     *
     * @throws \Exception
     *
     * @return string
     */
    public static function getContentLength($body)
    {
        return \count($body);
    }

    /**
     * @param bool   $condition
     * @param string $a
     * @param string $b
     *
     * @throws \Exception
     *
     * @return string
     */
    public static function _use($condition, $a, $b)
    {
        return $condition ? $a : $b;
    }

    /**
     * @param string $accessKeyId
     * @param string $accessKeySecret
     * @param string $versionPrefix
     *
     * @throws \Exception
     *
     * @return string
     */
    public static function getSignature($accessKeyId, $accessKeySecret, Request $request, $versionPrefix)
    {
        $queryForSign = [];
        if (0 === strpos($request->pathname, $versionPrefix . '/proxy/')) {
            $queryForSign = $request->query;
        }
        $auth = new Auth($accessKeyId, $accessKeySecret, '');

        return $auth->signRequest($request->method, $request->pathname, $request->headers, $queryForSign);
    }
}
