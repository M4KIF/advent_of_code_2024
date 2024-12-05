package org.runner;

import org.puzzle_1.PuzzleOne;
import org.puzzle_2.PuzzleTwo;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {

        var p1 = new PuzzleOne();
        var p2 = new PuzzleTwo();

        System.out.print("Result one: " + p1.solve_first() + "\n");
    }
}
