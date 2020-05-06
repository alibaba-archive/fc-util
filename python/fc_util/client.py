import hashlib
import base64
import hmac

from urllib import parse


def __build_canonical_headers(headers, prefix):
    list_sort = []
    fc_headers = {}
    for k in headers:
        lower_key = k.lower()
        if lower_key.startswith(prefix):
            list_sort.append(k)
            fc_headers[k] = headers.get(k)

    list_sort.sort()

    canonical = ""
    for k in list_sort:
        canonical += "{}:{}\n".format(k, fc_headers.get(k))

    return canonical


def __compose_string_to_sign(method, path, headers, queries):
    content_md5 = headers.get("content-md5") or ""
    content_type = headers.get("content-type") or ""
    date = headers.get("date") or ""
    sign_headers = __build_canonical_headers(headers, "x-fc-")

    pathname = parse.quote(path)
    str_to_sign = "{}\n{}\n{}\n{}\n{}{}".format(
        method, content_md5, content_type, date, sign_headers, pathname)

    if queries:
        sorted_keys = [k for k in queries]
        sorted_keys.sort()
        canonicalized_query_list = []
        for k in sorted_keys:
            canonicalized_query_list.append(
                parse.quote(k) + "=" + parse.quote(queries[k]))
        str_to_sign += "\n" + "\n".join(canonicalized_query_list)

    return str_to_sign


def get_content_md5(body):
    m = hashlib.md5()
    m.update(body)
    hash_bytes = m.digest()
    md5_str = str(base64.b64encode(hash_bytes), encoding="utf-8")
    return md5_str


def get_content_length(body):
    return len(body)


def use_(condition, a, b):
    return a if condition else b


def get_signature(access_key_id, access_key_secret, request, version_prefix):
    queries_to_sign = {}
    if request.pathname.startswith(version_prefix + "/proxy/"):
        queries_to_sign = request.query

    str_to_sign = __compose_string_to_sign(
        request.method, request.pathname, request.headers, queries_to_sign)
    digest_maker = hmac.new(bytes(access_key_secret, encoding="utf-8"),
                            bytes(str_to_sign, encoding="utf-8"), digestmod=hashlib.sha256)
    hash_bytes = digest_maker.digest()
    signed_str = str(base64.b64encode(hash_bytes), encoding="utf-8")
    return "FC {}:{}".format(access_key_id, signed_str)
