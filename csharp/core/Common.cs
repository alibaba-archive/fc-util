using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Web;

using AlibabaCloud.FCUtil.Utils;

using Tea;

namespace AlibabaCloud.FCUtil
{
    public static class Common
    {
        public static string GetContentMD5(byte[] bytes)
        {
            System.Security.Cryptography.MD5 md5 = new System.Security.Cryptography.MD5CryptoServiceProvider();
            byte[] data = md5.ComputeHash(bytes);
            return Convert.ToBase64String(data);
        }

        public static string GetContentLength(byte[] bytes)
        {
            return bytes.Length.ToString();
        }

        public static string GetSignature(string accessKeyId, string accessKeySecret, TeaRequest request, string versionPrefix)
        {
            Dictionary<string, string> queriesToSign = new Dictionary<string, string>();
            if (request.Pathname.StartsWith(versionPrefix + "/proxy/"))
            {
                queriesToSign = request.Query;
            }

            string stringToSign = ComposeStringToSign(request.Method, request.Pathname, request.Headers, queriesToSign);
            byte[] signData;
            using(KeyedHashAlgorithm algorithm = CryptoConfig.CreateFromName("HMACSHA256") as KeyedHashAlgorithm)
            {
                algorithm.Key = Encoding.UTF8.GetBytes(accessKeySecret);
                signData = algorithm.ComputeHash(Encoding.UTF8.GetBytes(stringToSign.ToCharArray()));
            }
            string signedStr = Convert.ToBase64String(signData);
            return string.Format("FC {0}:{1}", accessKeyId, signedStr);
        }

        public static string Use(bool condition, string a, string b)
        {
            return condition ? a : b;
        }

        public static bool Is4XXor5XX(int code)
        {
            return code >= 400 && code < 600;
        }

        internal static string ComposeStringToSign(string method, string path, Dictionary<string, string> headers, Dictionary<string, string> queries)
        {
            string contentMD5 = DictUtils.GetDicValue(headers, "content-md5").ToSafeString(string.Empty);
            string contentType = DictUtils.GetDicValue(headers, "content-type").ToSafeString(string.Empty);
            string date = DictUtils.GetDicValue(headers, "date");
            string signHeaders = BuildCanonicalHeaders(headers, "x-fc-");

            //Uri uri = new Uri(path);
            string pathName = HttpUtility.UrlDecode(path);
            string str = string.Format("{0}\n{1}\n{2}\n{3}\n{4}{5}", method, contentMD5, contentType, date, signHeaders, pathName);

            if (queries != null)
            {
                List<string> sortedKeys = queries.Keys.ToList();
                sortedKeys.Sort();
                StringBuilder canonicalizedQueryString = new StringBuilder();

                foreach (string key in sortedKeys)
                {
                    canonicalizedQueryString.Append("&")
                        .Append(PercentEncode(key)).Append("=")
                        .Append(PercentEncode(queries[key]));
                }
                str += "\n" + canonicalizedQueryString.ToString();
            }

            return str;
        }

        internal static string BuildCanonicalHeaders(Dictionary<string, string> headers, string prefix)
        {
            List<string> list = new List<string>();
            Dictionary<string, string> fcHeaders = new Dictionary<string, string>();

            foreach (var keypair in headers)
            {
                string lowerKey = keypair.Key.ToLower().Trim();
                if (lowerKey.StartsWith(prefix))
                {
                    list.Add(lowerKey);
                    fcHeaders[lowerKey] = DictUtils.GetDicValue(headers, keypair.Key);
                }
            }

            list.Sort();

            string canonical = string.Empty;
            foreach (string key in list)
            {
                canonical += string.Format("{0}:{1}\n", key, fcHeaders[key]);
            }

            return canonical;
        }

        internal static string PercentEncode(string value)
        {
            if (value == null)
            {
                return null;
            }
            var stringBuilder = new StringBuilder();
            var text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~";
            var bytes = Encoding.UTF8.GetBytes(value);
            foreach (char c in bytes)
            {
                if (text.IndexOf(c) >= 0)
                {
                    stringBuilder.Append(c);
                }
                else
                {
                    stringBuilder.Append("%").Append(string.Format(CultureInfo.InvariantCulture, "{0:X2}", (int) c));
                }
            }

            return stringBuilder.ToString().Replace("+", "%20")
                .Replace("*", "%2A").Replace("%7E", "~");
        }
    }
}
