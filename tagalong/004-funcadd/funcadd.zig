const std = @import("std");

pub fn add(a:i32, b:i32) i32 {
    return a+b;
}

pub fn main() void {
    std.debug.print("{}\n", .{add(42, 13)});
}