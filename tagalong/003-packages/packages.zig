const std = @import("std");
const RndGen = std.rand.DefaultPrng;

pub fn main() !void {
    var rnd = RndGen.init(0);
    var some_random_num = @mod(rnd.random().int(i32), 10);
    std.debug.print("my favorite number is {} and {}\n", .{some_random_num, std.math.pi});
}