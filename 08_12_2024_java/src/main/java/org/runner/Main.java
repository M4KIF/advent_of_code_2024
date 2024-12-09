package org.runner;

import org.src.Puzzle;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {

        Puzzle part_1 = new Puzzle("/data_final.txt");

        System.out.println("Part 1 result: " + part_1.SolvePart1());
        System.out.println("Part 2 results: " + part_1.SolvePart2());
    }
}
