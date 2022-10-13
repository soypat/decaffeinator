const std = @import("std");


pub fn main() void {
    var i:i32 = 0;
    var sum:i32 = 0;
    while (i < 10) {
        sum += i;
        i+=1;
    }
    std.debug.print("{}\n", .{sum});
}