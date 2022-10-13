const std = @import("std");

const Vertex = struct {
    x:i64,
    y:i64,
};

pub fn main() void {
    var v0 = Vertex{.x=0,.y=0};
    std.debug.print("{}\n", .{v0});

    var v1 = Vertex{.x=1,.y=2};
    std.debug.print("x: {}\n", .{v1.x});

    v1.x = 1e9;
    std.debug.print("new v1: {}\n", .{v1});
}