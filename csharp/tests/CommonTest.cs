using System.Collections.Generic;
using System.Text;

using AlibabaCloud.FCUtil;

using Tea;

using Xunit;

namespace tests
{
    public class CommonTest
    {
        [Fact]
        public void Test_GetContentMD5()
        {
            Assert.Equal("CY9rzUYh03PK3k6DJie09g==", Common.GetContentMD5(Encoding.UTF8.GetBytes("test")));
        }

        [Fact]
        public void Test_GetContentLength()
        {
            Assert.Equal("4", Common.GetContentLength(Encoding.UTF8.GetBytes("test")));
        }

        [Fact]
        public void Test_GetSignature()
        {
            TeaRequest request = new TeaRequest();
            request.Pathname = "/test";
            request.Method = "GET";
            string signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:tvvm2KiUDmXlQiuwz21CtN59JcKTT5QzIcrc122DKVo=", signedStr);

            Dictionary<string, string> header = new Dictionary<string, string>();
            header.Add("x-fc-test", "test");
            request.Headers = header;
            signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:lP6dXPCXlQNI7VY+f0brnZ4bwluNvGqcwdDf5GoM9ag=", signedStr);

            Dictionary<string, string> query = new Dictionary<string, string>();
            query.Add("testQuery", "test");
            request.Query = query;
            request.Pathname = "/version/proxy/";
            signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:aCD88kzuMQYKQGo7kq7DeqhT72nVrNSEqAJtiHNj9L0=", signedStr);
        }

        [Fact]
        public void Test_Use()
        {
            Assert.Equal("a", Common.Use(true, "a", "b"));
            Assert.Equal("b", Common.Use(false, "a", "b"));
            Assert.Equal("b", Common.Use(null, "a", "b"));
        }

        [Fact]
        public void Test_Is4XXor5XX()
        {
            Assert.True(Common.Is4XXor5XX(400));
            Assert.False(Common.Is4XXor5XX(200));
            Assert.False(Common.Is4XXor5XX(800));
        }
    }
}
