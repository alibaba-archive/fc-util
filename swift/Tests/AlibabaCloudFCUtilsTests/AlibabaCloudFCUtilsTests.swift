import XCTest
import Tea
@testable import AlibabaCloudFCUtils

final class AlibabaCloudFCUtilsTests: XCTestCase {
    func testGetContentMD5() {
        XCTAssertEqual("b45cffe084dd3d20d928bee85e7b0f21", AlibabaCloudFCUtils.getContentMD5("string".toBytes()))
    }

    func testGetContentLength() {
        XCTAssertEqual("6", AlibabaCloudFCUtils.getContentLength("string".toBytes()))
    }

    func testUse() {
        XCTAssertEqual("a", AlibabaCloudFCUtils.use(true, "a", "b"))
    }

    func testGetSignature() {
        let request: TeaRequest = TeaRequest()
        request.method = "GET"
        request.headers["content-md5"] = "b45cffe084dd3d20d928bee85e7b0f21"
        request.headers["content-type"] = "application/json"
        request.headers["date"] = "date"
        request.headers["x-fc-test"] = "test"

        XCTAssertEqual("077H9OLC6kXeHZWrZgsUAgOL5OkIL4/wh8DVOGB3+Rc=", AlibabaCloudFCUtils.getSignature("accessKeyId", accessKeySecret: "accessKeySecret", request, "x-fc-"))
    }

    static var allTests = [
        ("testGetContentMD5", testGetContentMD5),
    ]
}
