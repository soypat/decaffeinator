const std = @import("std");

var c:bool = false;
var python:bool = false;
var java:bool = false;
var s:i64 = 0;

var s2 = "This is long text";
var start:f64 = 1;
var stop:f64 = 20;

pub fn main() void {
    var afloat:f64 = 6.02;
    var casted:i64 = @floatToInt(i64, afloat);
    std.debug.print("{} {} {} [", .{c, python, java});

    // I just wanted to print the split array, 
    // this was the easiest way I could find.
    var words = std.mem.split(u8, s2, " ");
    while (words.next()) |chunk| {
        std.debug.print("{s}, ", .{chunk});
    }

    std.debug.print("] {} {} {} {}\n", .{start, stop, afloat, casted});
}