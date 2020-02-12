// This file is auto-generated, don't edit it
import * as url from 'url';
import { Readable } from 'stream';
import * as $tea from '@alicloud/tea-typescript';
import * as kitx from 'kitx';

function buildCanonicalHeaders(headers, prefix) {
  var list = [];
  var keys = Object.keys(headers);

  var fcHeaders = {};
  for (let i = 0; i < keys.length; i++) {
    let key = keys[i];

    var lowerKey = key.toLowerCase().trim();
    if (lowerKey.startsWith(prefix)) {
      list.push(lowerKey);
      fcHeaders[lowerKey] = headers[key];
    }
  }
  list.sort();

  var canonical = '';
  for (let i = 0; i < list.length; i++) {
    const key = list[i];
    canonical += `${key}:${fcHeaders[key]}\n`;
  }

  return canonical;
}

function composeStringToSign(method: string, path: string, headers, queries) {
  const contentMD5 = headers['content-md5'] || '';
  const contentType = headers['content-type'] || '';
  const date = headers['date'];
  const signHeaders = buildCanonicalHeaders(headers, 'x-fc-');

  const u = url.parse(path);
  const pathUnescaped = decodeURIComponent(u.pathname);
  var str = `${method}\n${contentMD5}\n${contentType}\n${date}\n${signHeaders}${pathUnescaped}`;

  if (queries) {
    var params = [];
    Object.keys(queries).forEach(function(key) {
      var values = queries[key];
      var type = typeof values;
      if (type === 'string') {
        params.push(`${key}=${values}`);
        return;
      }
      if (Array.isArray(values)) {
        queries[key].forEach(function(value){
          params.push(`${key}=${value}`);
        });
      }
    });
    params.sort();
    str += '\n' + params.join('\n');
  }
  return str;
}

export default class Client {

  static getContentMD5(body: Buffer): string {
    const digest = kitx.md5(body, 'hex');
    return Buffer.from(digest, 'utf8').toString('base64');
  }

  static getContentLength(body: Buffer): string {
    return '' + Buffer.byteLength(body);
  }

  static getSignature(accessKeyId: string, accessKeySecret: string, request: $tea.Request, versionPrefix: string): string {
    let queriesToSign = null;
    if (request.pathname.startsWith(`${versionPrefix}/proxy/`)) {
      queriesToSign = request.query;
    }
    const stringToSign = composeStringToSign(request.method, request.pathname, request.headers, queriesToSign);
    const sign = kitx.createHmac('sha256')(stringToSign, accessKeySecret, 'base64');
    return `FC ${accessKeyId}:${sign}`;
  }

  static use(condition: boolean, a: string, b: string): string {
    return condition ? a : b;
  }

  static is4XXor5XX(code: number): boolean {
    return code >= 400 && code < 600;
  }
}
