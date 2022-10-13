const std = @import("std");

const Allocator = std.mem.Allocator;

// I'm honestly flabberghasted the zig documentation has examples that do not run.
// this is one of them, copied from `allocator.zig`. I'm running 0.9.1, 12 october 2022.
pub fn sqrt(allocator: std.mem.Allocator, x:f64) []u8 {
    const result = try allocator.alloc(u8, 16);
    var fbs = std.io.fixedBufferStream(result);
    // Finding the fmt library was surprisingly hard. Zig's online presence has
    // a long way to go to be up to par to popular languages.
    if (x < 0) {
        std.fmt.format(fbs.writer(), "{f:16}i", std.math.sqrt(-x));
    } else {
        std.fmt.format(fbs.writer(), "{f:16}", std.math.sqrt(x));
    }
    return result;
}

pub fn main() void {
    var buffer: [100]u8 = undefined;
    const allocator = &std.heap.FixedBufferAllocator.init(&buffer).allocator;
    std.debug.print("{s} {s}\n", .{sqrt(allocator, 2), sqrt(allocator, -4)});
}