package org.puzzle_1;

import org.junit.jupiter.api.Test;
import org.src.Puzzle;

import java.util.List;

public class PuzzleTwoTest {

    @Test
    void TestIntegrationCalculateResonantAntinodePositionsDataSmall() {
        Puzzle testObj = new Puzzle("/data_3x.txt");
        System.out.println("Whole map: " + testObj.getMap());
        List<List<Integer>> points = testObj.RetrieveTowerLocations("X");

//        System.out.println("Locations: " + points);
//        Assertions.assertEquals(3, points.size());
//
//        List<List<List<Integer>>> subsets = testObj.CombinationsFromPointsArray(points);
//
//        System.out.println("Subsets created out of real data: " + subsets);
//
//        var result = testObj.CalculateResonantAntinodePositions(subsets);
//        System.out.println("Antinodes a: " + result);

        //Assertions.assertEquals(6, subsets.size());

        System.out.println(testObj.SolvePart2());
    }

}
