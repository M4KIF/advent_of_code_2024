package org.runner;

import org.src.Puzzle;

public class Main {
    public static void main(String[] args) throws InterruptedException {

        Puzzle part_1 = new Puzzle("/data.txt", 103, 101, 100);

        System.out.println("Part 1 result: " + part_1.Part1());
        System.out.println("Part 2 result: " + part_1.Part2());
    }
}
