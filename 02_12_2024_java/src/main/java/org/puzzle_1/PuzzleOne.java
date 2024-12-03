package org.puzzle_1;

import java.io.*;
import java.nio.charset.Charset;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Scanner;

public class PuzzleOne {

    private String data;

    public PuzzleOne() {
        try (var in = PuzzleOne.class.getResourceAsStream("/data.txt")) {
            this.data = new String(in.readAllBytes(), Charset.defaultCharset());
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public int solve_first() {

        //
        var safeSequences = 0;


        ArrayList<ArrayList<Integer>> data = getArrayLists();

        for(ArrayList<Integer> l: data) {

            var safe = true;
            var rising = false;
            var declining = false;

            for(int i = 1; i < l.size(); i++) {
                var temp = l.get(i-1) - l.get(i);

                if(i == 1) {
                    if (temp < 0) {
                        declining = true;
                    } else {
                        rising = true;
                    }
                }

                if (rising && temp < 0) {
                    safe = false;
                    break;
                }

                if (declining && temp > 0) {
                    safe = false;
                    break;
                }

                if(Math.abs(temp) < 1 || Math.abs(temp) > 3) {
                    safe = false;
                    break;
                }
            }

            if(safe) {
                safeSequences++;
            }
        }

        return safeSequences;
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
