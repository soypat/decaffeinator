const std = @import("std");

const result = struct {
    down: i64,
    up: i64,
};

pub fn collatz(a:i64) result {
    return result{
        .down = @divFloor(a, 2),
        .up   = a*3+1,
    };
}

pub fn main() void {
    const v:i64 = 60;
    var r = collatz(v);
    r.up = 2;
    std.debug.print("empezando en {} hay que saber subir {}, y bajar {}\n", .{v, r.up, r.down});
}