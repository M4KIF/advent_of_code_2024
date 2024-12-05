package org.puzzle_1;

import java.io.*;
import java.nio.charset.Charset;
import java.nio.file.Path;
import java.util.*;
import java.util.function.Function;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

import static java.util.stream.Collectors.counting;
import static java.util.stream.Collectors.groupingBy;

public class PuzzleOne {

    private String data;
    private List<List<Integer>> ruleSets;
    private List<Map<Integer, Integer>> updates;

    public PuzzleOne() {
        try (var in = PuzzleOne.class.getResourceAsStream("/data.txt")) {
            this.data = new String(in.readAllBytes(), Charset.defaultCharset());
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        this.ruleSets = new ArrayList<>();
        this.updates = new ArrayList<>();
    }

    public Set<List<Integer>> find_appliable_rules(List<Integer> values) {
        // Checking the rules which contain the given value on the left

            return this.ruleSets.stream()
                    .filter(a -> values.contains(a.getFirst()) && values.contains(a.getLast())).collect(Collectors.toSet());
//            return values.stream()
//                    .flatMap(v -> this.ruleSets.stream()
//                    .filter(a -> Objects.equals(a.getFirst(), v) && Objects.equals(a.getLast(), v)))
//                            .collect(Collectors.toSet());
    }

    public int check_right_side(List<List<Integer>> rules, int start, List<Integer> update)
    {
        for(int i = ++start; i < update.size(); i++) {
            int finalI = i;
            rules.stream().filter(a -> a.getLast()  == update.get(finalI)).toList();
        }
        return 0;
    }

    public int check_left_side() {
        return 0;
    }

    public int solve_first() {
        readData();

//        System.out.println(this.updates.get(0));
//        System.out.println(this.ruleSets);
        int jeden =0;

        for (Map<Integer, Integer> u : this.updates) {
            Set<List<Integer>> rules = find_appliable_rules(new ArrayList<>(u.keySet()));
            //System.out.println(u.keySet());

            int match_count = 0;
            for (List<Integer> r : rules) {
                //System.out.println(u);
                //System.out.println(r);
                if (u.containsKey(r.getFirst()) && u.containsKey(r.getLast())) {
                    if (u.get(r.getFirst()) < u.get(r.getLast())) {
                        match_count++;
                    }
                }
            }
            ;
            if (match_count == rules.size()) {
                Map<Integer, Integer> newMap = u.entrySet().stream().collect(Collectors.toMap(Map.Entry::getValue, Map.Entry::getKey));
                //System.out.println(newMap);
                jeden += newMap.get(newMap.size()/2);
            }
            //System.out.println("Maczow " + match_count + " ruli " + rules.size());
        }

        return jeden;
    }

    private void readData() {
        Scanner myReader = new Scanner(this.data);

        Pattern rulePattern = Pattern.compile("\\d+\\|\\d+");
        Pattern updateMatcher = Pattern.compile("[\\d+,]+[\\d+]");

        // Reading the data and creating the rulesets and updates lists
        while (myReader.hasNextLine()) {

            String line = myReader.nextLine();

            // Input validation for the peace of mind

            // For rule logic
            Matcher ruleLineMatcher = rulePattern.matcher(line);
            if (ruleLineMatcher.matches()) {
                var nums = line.split("\\|");
                this.ruleSets.add(Stream.of(nums).mapToInt(Integer::parseInt).boxed().collect(Collectors.toList()));
                continue;
            }

            // For update logic
            Matcher updateLineMatcher = updateMatcher.matcher(line);
            if (updateLineMatcher.matches()) {
                var nums = line.split(",");
                int count = 0;
                Map<Integer, Integer> map = IntStream.range(0, nums.length)
                        .boxed()
                        .collect(Collectors.toMap(
                                Function.identity(),
                                i -> Integer.parseInt(nums[i]))
                        );
                Map<Integer, Integer> newMap = map.entrySet().stream().collect(Collectors.toMap(Map.Entry::getValue, Map.Entry::getKey));
//                Map<Integer, Integer> newMap = new HashMap<>();
//
//                for(Map.Entry<Integer, Integer> entry : map.entrySet())
//                    newMap.put(entry.getValue(), entry.getKey());

                this.updates.add(newMap);
            }
        }
    }

}
