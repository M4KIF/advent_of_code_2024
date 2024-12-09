package org.puzzle_1;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions;
import java.lang.NullPointerException;
import java.util.ArrayList;
import java.util.List;

public class PuzzleOneTest {

    @Test
    void TestInstantiatingPuzzleOneWrongPath() {
        Assertions.assertThrows(NullPointerException.class, () -> {new PuzzleOne(".;'dgdsfg;'dsfds");});
    }

    @Test
    void TestInstantiatingPuzzleDataSmallValiditySizeCheck() {
        PuzzleOne testObj = new PuzzleOne("/data_small.txt");

        Assertions.assertEquals( 144, testObj.getMap().size());
    }

    @Test
    void TestInstantiatingPuzzleDataSmallValidityContentCheck() {
        PuzzleOne testObj = new PuzzleOne("/data_small.txt");

        ArrayList<Integer> knownZeroPos = new ArrayList<>();
        knownZeroPos.add(1);
        knownZeroPos.add(8);

        ArrayList<Integer> knownNotZeroPos = new ArrayList<>();
        knownNotZeroPos.add(1);
        knownNotZeroPos.add(9);

        Assertions.assertNotNull(testObj.getMap().get(knownZeroPos));
        System.out.println("Objects in the testObj at position (1,8): " + testObj.getMap().get(knownZeroPos));
        Assertions.assertTrue(testObj.getMap().get(knownZeroPos).contains("0"));

        Assertions.assertNotNull(testObj.getMap().get(knownNotZeroPos));
        System.out.println("Objects in the testObj at position (1,9): " + testObj.getMap().get(knownNotZeroPos));
        Assertions.assertFalse(testObj.getMap().get(knownNotZeroPos).contains("0"));
    }

    @Test
    void TestInstantiatingPuzzleDataFinalQuasiValiditySizeCheck() {
        PuzzleOne testObj = new PuzzleOne("/data_final.txt");

        System.out.println("Real size of the input: " + testObj.getMap().size());
        Assertions.assertNotEquals( 144, testObj.getMap().size());
    }

    @Test
    void TestRetrieveTowerLocations() {
        PuzzleOne testObj = new PuzzleOne("/data_small.txt");
        List<List<Integer>> points = testObj.RetrieveTowerLocations("A");

        Assertions.assertEquals(3, points.size());
    }

    @Test
    void TestCombinationsGeneration3Points() {
        // Containing the points
        List<List<Integer>> points = new ArrayList<>();

        // Examples
        List<Integer> point_1 = new ArrayList<>();
        point_1.add(4);
        point_1.add(8);

        List<Integer> point_2 = new ArrayList<>();
        point_2.add(3);
        point_2.add(6);

        List<Integer> point_3 = new ArrayList<>();
        point_3.add(7);
        point_3.add(2);

        // Creating the points array
        points.add(point_1);
        points.add(point_2);
        points.add(point_3);

        // Main test body
        PuzzleOne testObj = new PuzzleOne("");
        var result = testObj.CombinationsFromPointsArray(points);

        System.out.println("The combinations received: " + result);

        // Lazy asserion
        Assertions.assertNotNull(result);

        // This is the proper one, as the possible combinations without
        // reincluding the same elements is
        // n - number of points
        // n * (n-1)
        Assertions.assertEquals(6, result.size());
    }

    @Test
    void TestIntegrationRetrieveTowerLocationsAndCreateSubsets() {
        PuzzleOne testObj = new PuzzleOne("/data_small.txt");
        List<List<Integer>> points = testObj.RetrieveTowerLocations("A");

        Assertions.assertEquals(3, points.size());

        List<List<List<Integer>>> subsets = testObj.CombinationsFromPointsArray(points);

        System.out.println("Subsets created out of real data: " + subsets);

        Assertions.assertEquals(6, subsets.size());
    }

    @Test
    void TestPuzzleTowerDistanceCalc() {
        PuzzleOne testObj = new PuzzleOne("/data_small.txt");

        System.out.println(testObj.getMap());

        Assertions.assertEquals( 144, testObj.getMap().size());

        List<List<Integer>> points = testObj.RetrieveTowerLocations("A");

        Assertions.assertEquals(3, points.size());

        List<List<List<Integer>>> subsets = testObj.CombinationsFromPointsArray(points);

        var result = testObj.CalculateAntinodePositions(subsets);
        System.out.println(subsets);
        System.out.println("Antinodes A: " + result);

        List<List<Integer>> points1 = testObj.RetrieveTowerLocations("0");

        Assertions.assertEquals(3, points.size());

        List<List<List<Integer>>> subsets1 = testObj.CombinationsFromPointsArray(points1);

        var result1 = testObj.CalculateAntinodePositions(subsets1);
        System.out.println(subsets1);
        System.out.println("Antinodes 0: " + result1);

    }

    @Test
    void TestPuzzlePart1Solution() {
        PuzzleOne testObj = new PuzzleOne("/data_3a.txt");

        System.out.println(testObj.getMap());

        System.out.println("P1 result: " + testObj.SolvePart1());

    }

    @Test
    void TestIntegrationRetrieveTowerLocationsAndCreateSubsets3a() {
        PuzzleOne testObj = new PuzzleOne("/data_3a.txt");
        System.out.println("Whole map: " + testObj.getMap());
        List<List<Integer>> points = testObj.RetrieveTowerLocations("a");

        System.out.println("Locations: " + points);
        Assertions.assertEquals(3, points.size());

        List<List<List<Integer>>> subsets = testObj.CombinationsFromPointsArray(points);

        System.out.println("Subsets created out of real data: " + subsets);

        var result = testObj.CalculateAntinodePositions(subsets);
        System.out.println("Antinodes a: " + result);

        Assertions.assertEquals(6, subsets.size());
    }

}
