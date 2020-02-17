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
            request.Method = "Get";
            string signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:dPtao8SaRUV+WpLeN6SCT3wRcJrn8M6d9AoH8tpWj7k=", signedStr);

            Dictionary<string, string> header = new Dictionary<string, string>();
            header.Add("x-fc-test", "test");
            request.Headers = header;
            signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:mY3e3sR51Qu0+v/J8ObF819YniIezoEzmQhRo3iLBPo=", signedStr);

            Dictionary<string, string> query = new Dictionary<string, string>();
            query.Add("testQuery", "test");
            request.Query = query;
            request.Pathname = "/version/proxy/";
            signedStr = Common.GetSignature("accessKeyId", "accessKeySecret", request, "/version");
            Assert.Equal("FC accessKeyId:ljK3OKZ+fhhOqJu/PS5PV4BGvJzxdX/NGYc3NbNCqRs=", signedStr);
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
