package com.aliyun.fcutil;

import com.aliyun.tea.TeaRequest;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import javax.xml.bind.DatatypeConverter;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.security.InvalidKeyException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.*;

public class Common {
    public static String getContentMD5(byte[] body) throws NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("MD5");
        byte[] messageDigest = md.digest(body);
        return Base64.getEncoder().encodeToString(messageDigest);
    }

    public static String getContentLength(byte[] body) {
        if (null == body) {
            return "0";
        }
        return String.valueOf(body.length);
    }

    public static String use(boolean condition, String a, String b) {
        return condition ? a : b;
    }

    public static String getSignature(String accessKeyId, String accessKeySecret, TeaRequest request, String versionPrefix) throws UnsupportedEncodingException, NoSuchAlgorithmException, InvalidKeyException {
        Map<String, String> queriesToSign = null;
        String verify = versionPrefix + "/proxy/";
        if (request.pathname.startsWith(verify)) {
            queriesToSign = request.query;
        }
        String stringToSign = composeStringToSign(request.method, request.pathname, request.headers, queriesToSign);
        Mac mac = Mac.getInstance("HmacSHA256");
        mac.init(new SecretKeySpec(accessKeySecret.getBytes("UTF-8"), "HmacSHA256"));
        byte[] signData = mac.doFinal(stringToSign.getBytes("UTF-8"));
        String signedStr = DatatypeConverter.printBase64Binary(signData);
        return String.format("FC %s:%s", accessKeyId, signedStr);
    }

    public static boolean is4XXor5XX(Number code) {
        return code.intValue() >= 400 && code.intValue() < 600;
    }

    private static String buildCanonicalHeaders(Map<String, String> headers, String prefix) {
        List<String> list = new ArrayList<>();
        Set<String> keys = headers.keySet();
        Map<String, String> fcHeaders = new HashMap<>();
        String lowerKey;
        for (String key : keys) {
            lowerKey = key.toLowerCase().trim();
            if (lowerKey.startsWith(prefix)) {
                list.add(lowerKey);
                fcHeaders.put(key, headers.get(key));
            }
        }
        String[] sortedKeys = list.toArray(new String[]{});
        Arrays.sort(sortedKeys);

        StringBuilder stringBuilder = new StringBuilder();
        for (int i = 0; i < sortedKeys.length; i++) {
            String key = sortedKeys[i];
            stringBuilder.append(key).append(":").append(fcHeaders.get(key)).append("\n");
        }
        return stringBuilder.toString();
    }

    private static String composeStringToSign(String method, String path, Map<String, String> headers, Map<String, String> queries) throws UnsupportedEncodingException {
        String contentMD5 = null == headers.get("content-md5") ? "" : headers.get("content-md5");
        String contentType = null == headers.get("content-type") ? "" : headers.get("content-type");
        String date = null == headers.get("date") ? "" : headers.get("date");
        String signHeaders = buildCanonicalHeaders(headers, "x-fc-");
        String pathUnescaped = URLEncoder.encode(path, "UTF-8");
        String str = method + "\n" + contentMD5 + "\n" + contentType + "\n" + date + "\n" + signHeaders + pathUnescaped;
        if (null != queries) {
            List<String> params = new ArrayList<>();
            for (Map.Entry<String, String> entry : queries.entrySet()) {
                params.add(String.format("%s=%s", entry.getKey(), entry.getValue()));
            }
            String[] sortedKeys = params.toArray(new String[]{});
            Arrays.sort(sortedKeys);
            str += "\n" + String.join("\n", sortedKeys);
        }
        return str;
    }
}
