package com.aliyun.fcutil;

import com.aliyun.tea.TeaRequest;
import org.junit.Assert;
import org.junit.Test;

import java.util.HashMap;
import java.util.Map;

public class CommonTest {
    @Test
    public void getContentMD5Test() throws Exception {
        Assert.assertEquals("CY9rzUYh03PK3k6DJie09g==", Common.getContentMD5("test".getBytes("UTF-8")));
    }

    @Test
    public void getContentLengthTest() throws Exception{
        Assert.assertEquals("0", Common.getContentLength(null));
        Assert.assertEquals("5", Common.getContentLength("hello".getBytes("UTF-8")));
    }

    @Test
    public void getSignature() throws Exception{
        TeaRequest teaRequest =  new TeaRequest();
        teaRequest.method = "GET";
        teaRequest.pathname = "/test";
        String sign = Common.getSignature("accessKeyId", "accessKeySecret", teaRequest, "/version");
        Assert.assertEquals("FC accessKeyId:1Lb8sVc+U50Yhy9qmN3WOnHzl5v9Xf246eZRLlIa00o=", sign);

        Map<String, String> header =  new HashMap<>();
        header.put("x-fc-test", "test");
        teaRequest.headers = header;
        sign = Common.getSignature("accessKeyId", "accessKeySecret", teaRequest, "/version");
        Assert.assertEquals("FC accessKeyId:a18kQj1aBMu09LwfBFEWx74glqRvl6wMmllWUJlRnE0=", sign);

        Map<String, String> query =  new HashMap<>();
        query.put("testQuery", "test");
        teaRequest.query = query;
        teaRequest.pathname = "/test/proxy/";
        sign = Common.getSignature("accessKeyId", "accessKeySecret", teaRequest, "/test");
        Assert.assertEquals("FC accessKeyId:z6PVWj6+M4C5ATjpaiek4Sj9j6aiNTffDK6a56QKvgI=", sign);
    }

    @Test
    public void useTest() {
        Assert.assertEquals("a", Common.use(true, "a", "b"));
        Assert.assertEquals("b", Common.use(false, "a", "b"));
    }

    @Test
    public void is4XXor5XXTest() {
        Assert.assertTrue(Common.is4XXor5XX(400));
        Assert.assertFalse(Common.is4XXor5XX(200));
        Assert.assertFalse(Common.is4XXor5XX(800));
    }

}
