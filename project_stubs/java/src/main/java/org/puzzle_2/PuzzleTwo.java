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

    public boolean safe_sequence(ArrayList<Integer> data, int v_i) {
        this.recursive_calls++;

        if(v_i >= data.size()) {
            return false;
        }

        ArrayList<Integer> dataClone;
        if (v_i >= 0) {
            dataClone = (ArrayList<Integer>) data.clone();
            dataClone.remove(v_i);
        } else {
            dataClone = data;
        }

        return is_safe(dataClone) || safe_sequence(data, ++v_i);
    }

    public boolean is_safe(ArrayList<Integer> l) {
        Set<Integer> incr = Stream.of(1, 2, 3)
                .collect(Collectors.toCollection(HashSet::new));

        Set<Integer> decr = Stream.of(-1, -2, -3)
                .collect(Collectors.toCollection(HashSet::new));

        for(int i = 1; i < l.size(); i++) {
            var temp = l.get(i) - l.get(i-1);

            incr.add(temp);
            decr.add(temp);
        }

        return (incr.size() == 3 || decr.size() == 3);
    }

    /*
    What are the cases when the error can be recovered?
    - monotonicity change from rising to decreasing, then i can rollback from the current i and leave the n_i
    - from decreasing to rising
    - when the change is bigger than domain
     */
    public Boolean solve(ArrayList<Integer> d, int i) {
        if (i > d.size() - 1) {
            return false;
        }

        boolean result = safe(d);

        d.remove(i);
        return result ? result : solve(d, i);
    }

    public int solve_second() {

        //
        var safeSequences = 0;


        ArrayList<ArrayList<Integer>> data = getArrayLists();

        for(ArrayList<Integer> l: data) {

            if (safe_sequence(l, -1)) {
                safeSequences++;
            }
        }

        System.out.print("Recursive calls: " + this.recursive_calls + "\n");

        return safeSequences;
    }

    public int solve_one_and_a_half() {
        //
        var safeSequences = 0;


        ArrayList<ArrayList<Integer>> data = getArrayLists();

        // bylo 301

        for(ArrayList<Integer> l: data) {
            this.iterative_calls++;
            if(safe(l)) {
                safeSequences++;
            } else {
                for (int i = 0; i < l.size(); i++) {
                    this.iterative_calls++;
                    int removed = l.remove(i);
                    if (safe(l)) {
                        safeSequences++;
                        break;
                    }
                    l.add(i, removed);
                }
            }

        }

        System.out.print("Iterative calls: " + this.iterative_calls + "\n");

        return safeSequences;
    }

    public boolean safe(ArrayList<Integer> l) {

            Set<Integer> incr = Stream.of(1, 2, 3)
                    .collect(Collectors.toCollection(HashSet::new));

            Set<Integer> decr = Stream.of(-1, -2, -3)
                    .collect(Collectors.toCollection(HashSet::new));

            for(int i = 1; i < l.size(); i++) {
                var temp = l.get(i) - l.get(i-1);

                incr.add(temp);
                decr.add(temp);
            }

            return (incr.size() == 3 || decr.size() == 3);
    }

    private Boolean checkWithExclusions(ArrayList<Integer> l, int excluded) {
        ArrayList<Integer> copy = (ArrayList<Integer>) l.clone();
        copy.remove(excluded);

        var rising = false;
        var declining = false;

        var is_safe = true;
        for(int i = 1; i < copy.size();i++) {

            var temp = l.get(i-1) - l.get(i);

            if(i == 1) {
                if (temp < 0) {
                    declining = true;
                } else {
                    rising = true;
                }
                continue;
            }

            if (rising && temp < 0) {
                is_safe = false;
                break;
            }
            if (declining && temp > 0) {
                is_safe = false;
                break;
            }

            if(Math.abs(temp) < 1 || Math.abs(temp) > 3) {
                is_safe = false;
                break;
            }

        }

        return is_safe;
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
