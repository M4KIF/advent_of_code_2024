package org.runner;

import org.src.Puzzle;

public class Main {
    public static void main(String[] args) {

        Puzzle puzzle = new Puzzle("/data.txt");
        puzzle.loadDataFromResources();

        System.out.println("Part 1 result: " + puzzle.Part1(25, puzzle.getStones()));
        System.out.println("Part 2 result: " + puzzle.Part2(75, puzzle.getStones()));

        // Will It explode?
        System.out.println("Having some fun, result at 250: " + puzzle.Part2(250, puzzle.getStones()));
        System.out.println("Having some fun, result at 500: " + puzzle.Part2(500, puzzle.getStones()));
    }
}
