package org.puzzle_1;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions;
import org.src.Puzzle;

import java.lang.NullPointerException;
import java.math.BigInteger;
import java.util.ArrayList;

public class PuzzleTest {

    @Test
    void TestInstantiatingPuzzleOneWrongPath() {
        Assertions.assertThrows(NullPointerException.class, () -> {new Puzzle(".;'dgdsfg;'dsfds");});
    }

    @Test
    void TestPart1BlinkMethodWithTestData1() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();
        testObj.blink_part_1(testObj.getStones());
        Assertions.assertEquals("253000 1 7", testObj.getData());
    }

    @Test
    void TestSolvePart1MethodWithTestData11Blink() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        Assertions.assertEquals(3, testObj.Part1(1, testObj.getStones()));
    }

    @Test
    void TestSolvePart1MethodWithTestData16Blink() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        Assertions.assertEquals(22, testObj.Part1(6, testObj.getStones()));
    }

    @Test
    void TestSolvePart1MethodWithTestData41078Blink25() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(41078L));

        Assertions.assertEquals(41229, testObj.Part1(25, data));
    }

    @Test
    void TestSolvePart1MethodWithTestData8314Blink25() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(8314L));

        Assertions.assertEquals(47017, testObj.Part1(25, data));
    }

    @Test
    void TestSolvePart1MethodWithTestData125Blink25() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(125L));

        Assertions.assertEquals(19025, testObj.Part1(25, data));
    }

    @Test
    void TestSolvePart2MethodWithTestData125Blink25() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(125L));

        Assertions.assertEquals(BigInteger.valueOf(19025L), testObj.Part2(25, data));
    }

    @Test
    void TestSolvePart2MethodWithTestData8314Blink25() {
        Puzzle testObj = new Puzzle("src/test/resources/test/test_data_1.txt");
        testObj.loadData();

        ArrayList<BigInteger> data = new ArrayList<>();
        data.add(BigInteger.valueOf(8314L));

        Assertions.assertEquals(BigInteger.valueOf(47017L), testObj.Part2(25, data));
    }

}
