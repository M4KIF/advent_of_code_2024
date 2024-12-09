package org.runner;

import org.puzzle_1.PuzzleOne;
import org.puzzle_2.PuzzleTwo;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {

        PuzzleOne part_1 = new PuzzleOne("/data_final.txt");

        System.out.println("Part 1 result: " + part_1.SolvePart1());
    }
}
