
import hashlib
import base64
import hmac

from urllib import parse


class Client:
    @staticmethod
    def get_content_md5(body):
        m = hashlib.md5()
        m.update(body)
        hash_bytes = m.digest()
        md5_str = str(base64.b64encode(hash_bytes), encoding="utf-8")
        return md5_str

    @staticmethod
    def get_content_length(body):
        return body.__len__()

    @staticmethod
    def use(condition, a, b):
        return a if condition else b

    @staticmethod
    def get_signature(accessKeyId, accessKeySecret, request, versionPrefix):
        queries_to_sign = {}
        if request.pathname.startswith(versionPrefix + "/proxy/"):
            queries_to_sign = request.query

        str_to_sign = Client.__compose_string_to_sign(
            request.method, request.pathname, request.headers, queries_to_sign)
        print("GetSignature:stringToSign is " + str_to_sign)
        digest_maker = hmac.new(bytes(accessKeySecret, encoding="utf-8"),
                                bytes(str_to_sign, encoding="utf-8"), digestmod=hashlib.sha256)
        hash_bytes = digest_maker.digest()
        signed_str = str(base64.b64encode(hash_bytes), encoding="utf-8")
        return "FC {}:{}".format(accessKeyId, signed_str)

    @staticmethod
    def __compose_string_to_sign(method, path, headers, queries):
        content_md5 = headers.get("content-md5") or ""
        content_type = headers.get("content-type") or ""
        date = headers.get("date") or ""
        sign_headers = Client.__build_canonical_headers(headers, "x-fc-")

        pathname = parse.quote(path)
        str_to_sign = "{}\n{}\n{}\n{}\n{}{}".format(
            method, content_md5, content_type, date, sign_headers, pathname)

        if queries:
            sortedKeys = [k for k in queries]
            sortedKeys.sort()
            canonicalized_query_str = ""
            canonicalized_query_list = []
            for k in sortedKeys:
                canonicalized_query_list.append(
                    parse.quote(k) + "=" + parse.quote(queries[k]))
            str_to_sign += "\n" + "\n".join(canonicalized_query_list)

        return str_to_sign

    @staticmethod
    def __build_canonical_headers(headers, prefix):
        list_sort = []
        fcHeaders = {}
        for k in headers:
            lowerKey = k.lower()
            if lowerKey.startswith(prefix):
                list_sort.append(k)
                fcHeaders[k] = headers.get(k)

        list_sort.sort()

        canonical = ""
        for k in list_sort:
            canonical += "{}:{}\n".format(k, fcHeaders.get(k))

        return canonical
