import Foundation
import SwiftyJSON
import CryptoSwift
import Tea

public class AlibabaCloudFCUtils {
    public static func getContentMD5(_ body: [UInt8]) -> String {
        String(bytes: body, encoding: .utf8)?.md5() ?? ""
    }

    public static func getContentLength(_ body: [UInt8]) -> String {
        String(body.count)
    }

    public static func use(_ condition: Bool?, _ a: String, _ b: String) -> String {
        condition != nil && condition == true ? a : b
    }

    public static func getSignature(_ accessKeyId: String, accessKeySecret: String, _ request: TeaRequest, _ versionPrefix: String) -> String {
        var queriesToSign: [String: Any] = [String: Any]()
        if request.pathname.starts(with: versionPrefix) {
            queriesToSign = request.query
        }
        let strToSign: String = self.getStringToSign(request.method, request.pathname, request.headers, queriesToSign)
        let r: [UInt8] = try! HMAC(key: accessKeySecret, variant: .sha256).authenticate(strToSign.bytes)
        return r.toBase64() ?? "";
    }

    public static func Is4XXor5XX(_ code: Int) -> Bool {
        code >= 400 && code < 600;
    }

    private static func getStringToSign(_ method: String, _ pathname: String, _ headers: [String: String], _ query: [String: Any]) -> String {
        let contentMD5: String = headers["content-md5"] ?? ""
        let contentType: String = headers["content-type"] ?? ""
        let date: String = headers["date"] ?? ""
        let signHeaders: String = self.getCanonicalizedHeaders(headers, prefix: "x-fc-")
        let pathName: String = pathname.urlEncode()

        let tmp: [String] = [
            method,
            contentMD5,
            contentType,
            date,
            signHeaders,
            pathName
        ]

        var str = tmp.joined(separator: "\n")

        if (query.count > 0) {
            var keys: [String] = Array(query.keys)
            keys = keys.sorted()
            var canonicalizedQueryString: String = ""
            for key in keys {
                canonicalizedQueryString.append("&")
                canonicalizedQueryString.append(key.urlEncode())
                canonicalizedQueryString.append("=")
                canonicalizedQueryString.append("\(query[key] ?? "")")
            }
            str = str + "\n" + canonicalizedQueryString
        }
        return str
    }

    private static func getCanonicalizedResource(_ pathname: String, _ query: [String: Any]) -> String {
        if (query.count <= 0) {
            return pathname;
        }
        var keys: [String] = [String]();
        for (key, _) in query {
            keys.append(key);
        }
        keys.sort();
        var result: String = pathname + "?";
        for key in keys {
            result.append(key);
            result.append("=");
            result.append(query[key] as! String);
        }
        return result;
    }

    private static func getCanonicalizedHeaders(_ headers: [String: String], prefix: String) -> String {
        var canonicalizedKeys: [String] = [String]();
        for (key, _) in headers {
            if (key.hasPrefix(prefix)) {
                canonicalizedKeys.append(key);
            }
        }
        canonicalizedKeys.sort();
        var result: String = "";
        var n = 0;
        for key in canonicalizedKeys {
            if n != 0 {
                result.append(contentsOf: "\n");
            }
            result.append(contentsOf: key);
            result.append(contentsOf: ":");
            result.append(contentsOf: headers[key]?.trimmingCharacters(in: .whitespacesAndNewlines) ?? "");
            n += 1;
        }
        return result;
    }
}

extension String {
    func urlEncode() -> String {
        let unreserved = "*-._"
        let allowedCharacterSet = NSMutableCharacterSet.alphanumeric()
        allowedCharacterSet.addCharacters(in: unreserved)
        allowedCharacterSet.addCharacters(in: " ")
        var encoded = addingPercentEncoding(withAllowedCharacters: allowedCharacterSet as CharacterSet)
        encoded = encoded?.replacingOccurrences(of: " ", with: "%20")
        return encoded ?? ""
    }

    func toBytes() -> [UInt8] {
        [UInt8](self.utf8)
    }
}