const std = @import("std");

const Allocator = std.mem.Allocator;

pub fn sqrt(allocator: Allocator, x:f64) ![]u8 {
    var result = try allocator.alloc(u8, 40);
    if (x < 0) {
        result = try std.fmt.bufPrint(result, "{}", .{std.math.sqrt(-x)});
     } else {
        result = try std.fmt.bufPrint(result, "{}", .{std.math.sqrt(x)});
     }
    return result;
}

pub fn main() void {
    var buffer: [100]u8 = undefined;
    const allocator = std.heap.FixedBufferAllocator.init(&buffer).allocator();
    std.debug.print("{s} {s}\n", .{sqrt(allocator, 2), sqrt(allocator, -4)});
}