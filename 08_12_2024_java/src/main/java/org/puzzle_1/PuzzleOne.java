package org.puzzle_1;

import lombok.Getter;
import lombok.Setter;

import java.io.*;
import java.nio.charset.Charset;
import java.util.*;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@Getter
@Setter
public class PuzzleOne {

    // Used path
    private String path;

    // Point - station/antinode map
    private Map<ArrayList<Integer>, ArrayList<String>> map;

    static class PointExcludor implements Comparator<List<Integer>> {
        @Override
        public int compare(List<Integer> o1, List<Integer> o2) {
            // Excluding when placed directly to the side
            if ((Math.abs(o1.getFirst() - o2.getFirst()) == 1 &&
                    Math.abs(o1.getLast() - o2.getLast()) == 0)
                    || (Math.abs(o1.getFirst() - o2.getFirst()) == 0 &&
                            Math.abs(o1.getLast() - o2.getLast()) == 1) ||
                    (Math.abs(o1.getFirst() - o2.getFirst()) == 0 &&
                            Math.abs(o1.getLast() - o2.getLast()) == 0)) {
                return 0;
            }
            return 1;
        }
    }

    public PuzzleOne(String filePath) {
        this.path = filePath;
        this.map = new HashMap<>();

        // Reading the input
        try (var in = PuzzleOne.class.getResourceAsStream(this.path)) {
            // For easier map creation
            AtomicInteger row = new AtomicInteger(0);
            AtomicInteger col = new AtomicInteger(0);

            String data = new String(in.readAllBytes(), Charset.defaultCharset());
            Scanner myReader = new Scanner(data);

            // Reading the data and calculating from the get go
            while (myReader.hasNextLine()) {
                String line = myReader.nextLine();
                String[] split = line.split("");
                System.out.println(split.length);

                Map<ArrayList<Integer>, ArrayList<String>> lineData =
                        Stream.of(split).map(s -> {
                            // Create the list that will hold all values
                            // The antenna id, antinodes etc.
                            ArrayList<String> entries = new ArrayList<>();
                            entries.add(s);

                            // Creating the pos list, would've preddered
                            // something more straight forward like std::Pair
                            // or a tuple, but whatever
                            ArrayList<Integer> pos = new ArrayList<>();
                            pos.add(row.get());
                            pos.add(col.get());

                            // Creating an entry that will soon get picked out in the collector
                                    // in order to be added to the real map
                            AbstractMap.SimpleEntry<ArrayList<Integer>, ArrayList<String>> crd_station_line
                            = new AbstractMap.SimpleEntry<>(pos, entries);
                            col.getAndIncrement();
                            return crd_station_line;})
                        .collect(Collectors.toMap(e-> e.getKey(),
                                e -> e.getValue()));

                row.getAndIncrement();
                col.getAndSet(0);
                lineData.entrySet().
                        stream().
                        forEachOrdered(e -> {
                            map.put(e.getKey(), e.getValue());
                        });
            }

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public List<String> RetrievePossibleFrequencyList() {
        return this.map.values().stream().
                flatMap(Collection::stream).
                filter(s -> !Objects.equals(s, "#") && !Objects.equals(s, ".")).
                collect(Collectors.toList());
    }

    public List<List<Integer>> RetrieveTowerLocations(String frequency) {
        return this.map.entrySet().stream().
                filter(e -> e.getValue().contains(frequency)).
                map(e -> e.getKey()).
                collect(Collectors.toList());
    }

    private void combinationsHelper(
            List<List<List<Integer>>> combinations,
            List<List<Integer>> points,
            List<List<Integer>> current) {

        // Base case if we'have a pair of points
        if (current.size() == 2) {
            List<List<Integer>> temp = new ArrayList<>(current);
            combinations.add(temp);
            return;
        }

        for(int i = 0; i < points.size(); i++) {
            if (current.contains(points.get(i))) {
                continue;
            }

            // Append to the current
            current.add(points.get(i));

            // Run recursion
            combinationsHelper(combinations, points, current);

            // Remove the previously added
            current.removeLast();
        }
    }

    public List<List<List<Integer>>> CombinationsFromPointsArray(List<List<Integer>> points) {
        List<List<List<Integer>>> combinations = new ArrayList<>();
        List<List<Integer>> current = new ArrayList<>();

        combinationsHelper(combinations, points, current);

        return combinations;
    }

    public List<Integer> CompareAndModifyPoints(List<Integer> first, List<Integer> second) {

        // Values that will get used for modification of the coordinates
        var x_delta = Math.abs(first.getLast() - second.getLast());
        var y_delta = Math.abs(first.getFirst() - second.getFirst());

        // New point containing an antinode
        List<Integer> antinode = new ArrayList<>();

        // Comparing the X coordinates
        var x_first = first.getLast();
        var x_second = second.getLast();

        if (x_first < x_second) {
            x_delta *= -1;
        }

        // Comparing the Y coordinates
        // also, the deciding step
        // on which point to modify
        var y_first = first.getFirst();
        var y_second = second.getFirst();

        if (y_first < y_second) {
            y_delta *= -1;
        }

        antinode.add(y_first + y_delta);
        antinode.add(x_first + x_delta);

        System.out.println("First point: " + first + " second: " + second + " result: " + antinode);
        return antinode;
    }

    public Set<List<Integer>> CalculateAntinodePositions(List<List<List<Integer>>> pairs) {

        PointExcludor excluding = new PointExcludor();
        return pairs.stream().
                filter(p -> { return excluding.compare(p.getFirst(), p.getLast()) != 0; }).
                map(p -> CompareAndModifyPoints(p.getFirst(), p.getLast())).
                filter(p -> this.map.containsKey(p)).
                collect(Collectors.toSet());
    }

    public int SolvePart1() {
        AtomicInteger result = new AtomicInteger(0);

        // Retrieving all of the possible frequencies
        RetrievePossibleFrequencyList().stream().
                map(frequency -> RetrieveTowerLocations(frequency)).
                map(locations -> CombinationsFromPointsArray(locations)).
                map(pairs -> CalculateAntinodePositions(pairs)).
                forEachOrdered(partial_result -> partial_result.stream().
                        forEachOrdered(entry -> this.map.get(entry).add("#")));

        System.out.println(this.map);
        return this.map.entrySet().stream().
                filter(e -> e.getValue().contains("#")).
                collect(Collectors.toList()).size();
    }
}
