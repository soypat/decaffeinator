const std = @import("std");

pub fn main() void {
    var a = [2][]const u8 {"", ""};
    a[0] = "Hello";
    a[1] = "World";
    std.debug.print("{s} {s}\n", .{a[0], a[1]});
    std.debug.print("{s}\n", .{a});

    var primes = [_]i64 {2, 3, 5, 7, 11, 13};
    std.debug.print("{any}\n", .{primes});
}