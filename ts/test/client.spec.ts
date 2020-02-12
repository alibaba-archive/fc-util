import Client from '../src/client';

import * as $tea from '@alicloud/tea-typescript';
import assert from 'assert';
import 'mocha';

describe('FC Util', function() {
    it('Module should ok', function() {
        assert.ok(Client);
    });

    it('getContentMD5', function () {
        assert.deepStrictEqual(Client.getContentMD5(Buffer.from('Hello world!')), 'ODZmYjI2OWQxOTBkMmM4NWY2ZTA0NjhjZWNhNDJhMjA=');
    });

    it('getContentLength', function () {
        assert.deepStrictEqual(Client.getContentLength(Buffer.from('Hello world!')), '12');
    });

    it('use', function() {
        assert.deepStrictEqual(Client.use(true, 'a', 'b'), 'a');
        assert.deepStrictEqual(Client.use(false, 'a', 'b'), 'b');
    });

    it('is4XXor5XX', function() {
        assert.deepStrictEqual(Client.is4XXor5XX(300), false);
        assert.deepStrictEqual(Client.is4XXor5XX(500), true);
        assert.deepStrictEqual(Client.is4XXor5XX(400), true);
    });

    describe('getSignature', function() {
        it('normal request', function () {
            const request = new $tea.Request();
            request.method = 'GET';
            request.pathname = '/';
            let sign = Client.getSignature('accessKeyId', 'accessKeySecret', request, '/version');
            assert.deepStrictEqual(sign, 'FC accessKeyId:oCcDVSb6OZJgygYpmvBCq73DHbox6djMYSa9KBBPRbU=');
        });

        it('with x-fc- prefix headers', function () {
            const request = new $tea.Request();
            request.method = 'GET';
            request.pathname = '/';
            request.headers = {
                'x-fc-id': 'id'
            };
            let sign = Client.getSignature('accessKeyId', 'accessKeySecret', request, '/version');
            assert.deepStrictEqual(sign, 'FC accessKeyId:OVypZ041SMFDNYevxvsIKtZ8ePFCMRVgII25vnEEhuI=');
        });

        it('with query', function () {
            const request = new $tea.Request();
            request.method = 'GET';
            request.pathname = '/version/proxy/ends';
            request.query = {
                'key': 'value'
            };
            request.headers = {
                'x-fc-id': 'id'
            };
            let sign = Client.getSignature('accessKeyId', 'accessKeySecret', request, '/version');
            assert.deepStrictEqual(sign, 'FC accessKeyId:8BMNbZlME8C9Vi9Ed6kZBB58y0O8IbWAw0T7Ad+Cv8Q=');
        });
    });
});
