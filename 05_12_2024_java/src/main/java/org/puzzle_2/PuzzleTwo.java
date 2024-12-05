package org.puzzle_2;

import org.puzzle_1.PuzzleOne;

import java.io.IOException;
import java.nio.charset.Charset;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class PuzzleTwo {

    private int recursive_calls = 0;
    private int iterative_calls = 0;
    private String data;

    public PuzzleTwo() {
        try (var in = PuzzleOne.class.getResourceAsStream("/data.txt")) {
            this.data = new String(in.readAllBytes(), Charset.defaultCharset());
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private ArrayList<ArrayList<Integer>> getArrayLists() {
        Scanner myReader = new Scanner(this.data);
        ArrayList<ArrayList<Integer>> data = new ArrayList<>();

        // Reading the data and calculating from the get go
        while (myReader.hasNextLine()) {
            String line = myReader.nextLine();

            ArrayList<Integer> lineList = new ArrayList<>();

            // Another reader for a string line
            Scanner lineReader = new Scanner(line);
            while(lineReader.hasNextInt()) {
                lineList.add(lineReader.nextInt());
            }

            data.add(lineList);
        }
        return data;
    }
}
