<?php

namespace AlibabaCloud\Tea\FCUtils\Tests;

use AlibabaCloud\Tea\FCUtils\FCUtils;
use AlibabaCloud\Tea\Request;
use PHPUnit\Framework\TestCase;

/**
 * @internal
 * @coversNothing
 */
class FCUtilsTest extends TestCase
{
    public function testGetContentMD5()
    {
        $content = 'this is content for test.';
        $val     = unpack('C*', $content);
        $this->assertEquals('5853d4df3b9a9b2682bf1d5b85e29b4a', FCUtils::getContentMD5($val));
    }

    public function testGetContentLength()
    {
        $content = 'this is content for test.';
        $val     = unpack('C*', $content);
        $this->assertEquals(25, FCUtils::getContentLength($val));
    }

    public function testUse()
    {
        $this->assertEquals('a', FCUtils::_use(true, 'a', 'b'));
        $this->assertEquals('b', FCUtils::_use(null, 'a', 'b'));
    }

    public function testGetSignature()
    {
        $request           = new Request();
        $request->pathname = '/test';
        $request->method   = 'GET';
        $sign              = FCUtils::getSignature('accessKeyId', 'accessKeySecret', $request, '/version');
        $this->assertEquals('FC accessKeyId:tvvm2KiUDmXlQiuwz21CtN59JcKTT5QzIcrc122DKVo=', $sign);

        $request->headers['x-fc-test'] = 'test';
        $sign                          = FCUtils::getSignature('accessKeyId', 'accessKeySecret', $request, '/version');
        $this->assertEquals('FC accessKeyId:lP6dXPCXlQNI7VY+f0brnZ4bwluNvGqcwdDf5GoM9ag=', $sign);

        $request->query['testQuery'] = 'test';
        $request->pathname           = '/version/proxy/';
        $sign                        = FCUtils::getSignature('accessKeyId', 'accessKeySecret', $request, '/version');
        $this->assertEquals('FC accessKeyId:aCD88kzuMQYKQGo7kq7DeqhT72nVrNSEqAJtiHNj9L0=', $sign);
    }
}
