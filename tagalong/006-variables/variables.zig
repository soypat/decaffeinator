const std = @import("std");

pub fn main() void {
    var s = "";
    var s2 = "Hello!";
   
    var i:i64 = 0; 
    var j:i64 = 42;
    var k:i64 = 12;
    std.debug.print("\"{s}\" \"{s}\" {} {} {}\n", .{s, s2, i, j, k});
}