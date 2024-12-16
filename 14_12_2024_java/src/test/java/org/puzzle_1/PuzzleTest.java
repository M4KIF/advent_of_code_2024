package org.puzzle_1;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions;
import org.src.Puzzle;

import java.awt.*;
import java.lang.NullPointerException;
import java.util.AbstractMap;

public class PuzzleTest {

    @Test
    void TestInstantiatingPuzzleOneWrongPath() {
        Assertions.assertThrows(NullPointerException.class, () -> {new Puzzle(".;'dgdsfg;'dsfds", 7, 11, 100);});
    }

    @Test
    void TestInstantiatingPuzzleValidData12() {
        Puzzle testObj = new Puzzle("/data_1.txt", 7, 11, 5);

        Assertions.assertEquals( 12, testObj.getRobots().size());
    }

    @Test
    void TestInstantiatingPuzzle1RobotWalk5secs() {
        Puzzle testObj = new Puzzle("/data_2.txt", 7, 11, 5);

        // retrieving the robot for testing purposes
        Puzzle.Robot testRobot = testObj.getRobots().getFirst();

        Assertions.assertEquals( new AbstractMap.SimpleEntry<>(1,4), testRobot.walk());
        Assertions.assertEquals( new AbstractMap.SimpleEntry<>(5,6), testRobot.walk());
        Assertions.assertEquals( new AbstractMap.SimpleEntry<>(2,8), testRobot.walk());
        Assertions.assertEquals( new AbstractMap.SimpleEntry<>(6,10), testRobot.walk());
        Assertions.assertEquals( new AbstractMap.SimpleEntry<>(3,1), testRobot.walk());
    }

    @Test
    void TestInstantiatingPuzzle112RobotWalk100secsPart1Solve() {
        Puzzle testObj = new Puzzle("/data_1.txt", 7, 11, 100);

        // retrieving the robot for testing purposes
        Assertions.assertEquals(12, testObj.Part1());
    }

    @Test
    void TestInstantiatingPuzzle1FinalRobotWalk100secsPart1Solve() {
        Puzzle testObj = new Puzzle("/data_3.txt", 103, 101, 100);

        // retrieving the robot for testing purposes
        Assertions.assertEquals(12, testObj.Part1());
    }

}
