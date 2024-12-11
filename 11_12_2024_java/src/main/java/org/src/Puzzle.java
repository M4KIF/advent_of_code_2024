package org.src;

import lombok.Getter;
import lombok.Setter;

import java.io.*;
import java.math.BigInteger;
import java.nio.charset.Charset;
import java.util.*;

@Getter
@Setter
public class Puzzle {

    // Used path
    private String path;
    private String data;
    private ArrayList<BigInteger> stones;

    private HashMap<BigInteger, BigInteger> memo;

    private BigInteger gowno;


    public Puzzle(String filePath) {
        this.path = filePath;
        this.stones = new ArrayList<>();
    }

    public void loadDataFromResources() {

        // Reading the input
        try (var in = Puzzle.class.getResourceAsStream(this.path)) {
            this.data = new String(in.readAllBytes(), Charset.defaultCharset());

            var splitStones = this.data.split(" ");

            for(String stone : splitStones) {
                this.stones.add(BigInteger.valueOf(Long.valueOf(stone)));
            }

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    // For testing only, redundancies etc.
    public void loadData() {

        BufferedReader reader = null;
        try {
            reader = new BufferedReader(new FileReader(this.path));
            this.data = reader.readLine();

            var splitStones = this.data.split(" ");

            for(String stone : splitStones) {
                this.stones.add(BigInteger.valueOf(Long.valueOf(stone)));
            }

            reader.close();
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    // Yeah right!? Who would've thought, learnt another exercise pattern today xD
    public ArrayList<BigInteger> blink_part_1(ArrayList<BigInteger> stones) {
        ArrayList<BigInteger> newStones = new ArrayList<>();

        for(BigInteger stone : stones) {
            var stoneString = stone.toString();
            if (stoneString.equals("0")) {
                // append with first rule
                newStones.add(BigInteger.valueOf(Long.valueOf("1")));
            } else if (stoneString.length() % 2 == 0) {
                // Append with second rule
                var firstStone = stoneString.substring(0, stoneString.length() / 2);
                var secondStone = stoneString.substring(stoneString.length() / 2);
                secondStone = String.valueOf(Integer.parseInt(secondStone));

                newStones.add(BigInteger.valueOf(Long.valueOf(firstStone)));
                newStones.add(BigInteger.valueOf(Long.valueOf(secondStone)));
            } else {
                // Append with third rule
                Long thirdRule = Long.parseLong(stoneString) * 2024;
                newStones.add(BigInteger.valueOf(Long.valueOf(String.valueOf(thirdRule))));

            }
        }

        return newStones;
    }

    // DP approach, ie. recursion dfs + memoisation
    public BigInteger blink_part_2(BigInteger number, HashMap<AbstractMap.SimpleEntry<BigInteger, Integer>, BigInteger> memo, int blinks) {
        var stoneString = number.toString();
        if(blinks == 0) {
            return BigInteger.ONE;
        }

        if (stoneString.equals("0")) {
            // First rule
            var firstRule = BigInteger.valueOf(1L);
            BigInteger res1;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(firstRule, blinks - 1))) {
                res1 = memo.get(new AbstractMap.SimpleEntry<>(firstRule, blinks - 1));
            } else {
                res1 = blink_part_2(firstRule, memo, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(firstRule, blinks - 1), res1);
            }
            return res1;
        } else if (stoneString.length() % 2 == 0) {
            // Second rule
            var secondRuleFirstStone = BigInteger.
                    valueOf(
                            Long.parseLong(
                                    stoneString.substring(0, stoneString.length() / 2)
                            )
                    );
            var secondRuleSecondStone = BigInteger.
                    valueOf(
                            Long.parseLong(
                                    String.valueOf(
                                            Integer.parseInt(
                                                    stoneString.substring(stoneString.length() / 2)
                                            )
                                    )
                            )
                    );

            BigInteger res1;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(secondRuleFirstStone, blinks - 1))) {
                res1 = memo.get(new AbstractMap.SimpleEntry<>(secondRuleFirstStone, blinks - 1));
            } else {
                res1 = blink_part_2(secondRuleFirstStone, memo, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(secondRuleFirstStone, blinks - 1), res1);
            }

            BigInteger res2;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(secondRuleSecondStone, blinks - 1))) {
                res2 = memo.get(new AbstractMap.SimpleEntry<>(secondRuleSecondStone, blinks - 1));
            } else {
                res2 = blink_part_2(secondRuleSecondStone, memo, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(secondRuleSecondStone, blinks - 1), res2);
            }

            return res1.add(res2);
        } else {
            var thirdRule = number.multiply(BigInteger.valueOf(2024L));
            BigInteger res2;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(thirdRule, blinks - 1))) {
                res2 = memo.get(new AbstractMap.SimpleEntry<>(thirdRule, blinks - 1));
            } else {
                res2 = blink_part_2(thirdRule, memo, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(thirdRule, blinks - 1), res2);
            }

            return res2;
        }

    }

    public int Part1(int times, ArrayList<BigInteger> stones) {

        for(int i = 0; i < times; i++) {
            stones = blink_part_1(stones);
        }

        return stones.size();
    }

    public BigInteger Part2(int times, ArrayList<BigInteger> stones) {
        // Keeping the result
        BigInteger result = BigInteger.ZERO;

        // Initialising the memo
        HashMap<AbstractMap.SimpleEntry<BigInteger, Integer>, BigInteger> memo = new HashMap<>();

        for(BigInteger stone : stones) {
            var temp = blink_part_2(stone, memo, times);
            result = result.add(temp);
        }

        return result;
    }


}
