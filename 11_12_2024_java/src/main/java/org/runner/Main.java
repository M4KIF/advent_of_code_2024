package org.runner;

import org.src.Puzzle;

import java.math.BigInteger;
import java.util.ArrayList;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {

        Puzzle part_1 = new Puzzle("/data.txt");
        part_1.loadDataFromResources();

        // Will It explode?
        //System.out.println("Part 1 result: " + part_1.Part1(25, part_1.getStones()));

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(Long.valueOf("125")));
        System.out.println("Part 2 result: " + part_1.Part2(75, part_1.getStones()));
    }
}
