package org.src;

import lombok.Getter;
import lombok.Setter;

import java.io.*;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.nio.charset.Charset;
import java.util.*;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.Collectors;
import java.util.stream.Stream;

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
        this.memo = new HashMap<>();
        this.gowno = BigInteger.ZERO;
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

//    public void blink_part_2(BigInteger number, BigInteger score,  int blinks) {
////        if(blinks == 0) {
//////            if (this.memo.containsKey(number)) {
//////                return this.memo.get(number);
//////            }
////            return BigInteger.ONE;
////        }
////
////        if (this.memo.containsKey(number)) {
////            return score.add(this.memo.get(number));
////        }
////
////        var stoneString = number.toString();
////
////        if (stoneString.length() % 2 == 0) {
////            // Append with second rule
////            var firstStone = stoneString.substring(0, stoneString.length() / 2);
////            var secondStone = stoneString.substring(stoneString.length() / 2);
////            secondStone = String.valueOf(Integer.parseInt(secondStone));
////
////            this.memo.put(BigInteger.valueOf(Long.valueOf(firstStone)), blink_part_2(BigInteger.valueOf(Long.valueOf(firstStone)), BigInteger.ZERO,blinks - 1));
////            this.memo.put(BigInteger.valueOf(Long.valueOf(secondStone)), blink_part_2(BigInteger.valueOf(Long.valueOf(secondStone)), BigInteger.ZERO, blinks - 1));
////
////            score = score.add(memo.get(BigInteger.valueOf(Long.valueOf(firstStone))));
////            score = score.add(memo.get(BigInteger.valueOf(Long.valueOf(secondStone))));
////        } else if (!number.equals(BigInteger.ZERO)) {
////            // Append with third rule
////            this.memo.put(number, blink_part_2(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), score, --blinks));
////            score = score.add(this.memo.get(number));
////
////        }
////        return score;
//
//        // Base case
//        if(blinks <= 0) {
//            this.memo.put(number, BigInteger.ONE);
//            this.gowno = this.gowno.add(BigInteger.ONE);
//            System.out.println("End - Number: " + number + " score: " + this.gowno + " blink: " + blinks);
//
//            return;
//        }
//
//        if (this.memo.containsKey(number)) {
//            //this.memo.put(number, this.memo.get(number).multiply(BigInteger.TWO));
//            System.out.println("Memo - Number: " + number + " score: " + this.memo.get(number) + " blink: " + blinks);
//            this.gowno = this.gowno.add(this.memo.get(number));
//            return;
//        }
//
//        var stoneString = number.toString();
//        if (stoneString.equals("0")) {
//            this.memo.put(number, BigInteger.ONE);
//            System.out.println("Rule one - Number: " + number + " score: " + this.gowno + " blink: " + blinks);
//            this.gowno = this.gowno.add(BigInteger.ONE);
//
//        } else if (stoneString.length() % 2 == 0) {
//            // Append with second rule
//            var firstStone = stoneString.substring(0, stoneString.length() / 2);
//            var secondStone = stoneString.substring(stoneString.length() / 2);
//            secondStone = String.valueOf(Integer.parseInt(secondStone));
//
//            blink_part_2(BigInteger.valueOf(Long.valueOf(firstStone)), this.gowno, blinks - 1);
//
//            blink_part_2(BigInteger.valueOf(Long.parseLong(secondStone)), this.gowno, blinks - 1);
//
//            this.memo.put(number, this.gowno);
//            System.out.println("Rule two - Number: " + number + " score: " + this.gowno + " blink: " + blinks);
//            this.gowno = this.gowno.add(BigInteger.ONE);
//
//
//        } else {
//            // Append with third rule
//            blink_part_2(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), this.gowno, blinks - 1);
//
//            this.memo.put(number, this.gowno);
//            System.out.println("Rule third - Number: " + number + " score: " + this.gowno + " blink: " + blinks);
//            this.gowno = this.gowno.add(BigInteger.ONE);
//
//        }
//    }
    public BigInteger blink_part_2(BigInteger number, HashMap<AbstractMap.SimpleEntry<BigInteger, Integer>, BigInteger> memo, BigInteger score, int blinks) {
        var stoneString = number.toString();
        if(blinks == 0) {
            //System.out.println("End - Number: " + number + " blink: " + blinks+ " score: " + score);
//            if (!memo.containsKey(number) && values.containsKey(number.toString())) {
//                memo.put(number, values.get(number.toString()));
//            }
            return BigInteger.ONE;
        }

        if (stoneString.equals("0")) {
            //System.out.println("Rule one - Number: " + number + " blink: " + blinks+ " score: " + score);
            BigInteger res1;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong("1")), blinks - 1))) {
                res1 = memo.get(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong("1")), blinks - 1));
            } else {
                res1 = blink_part_2(BigInteger.valueOf(Long.parseLong("1")), memo, score, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong("1")), blinks - 1), res1);
            }
//            if (!memo.containsKey(BigInteger.valueOf(Long.parseLong("1")))) {
//                values.put(stoneString, res1);
//            }
            score = score.add(res1);
        } else if (stoneString.length() % 2 == 0) {
            // Append with second rule
            var firstStone = stoneString.substring(0, stoneString.length() / 2);
            var secondStone = stoneString.substring(stoneString.length() / 2);
            secondStone = String.valueOf(Integer.parseInt(secondStone));

            //var res1 = blink_part_2(BigInteger.valueOf(Long.parseLong(firstStone)), values, memo, score, blinks - 1);

//            BigInteger res1;
//            if (memo.containsKey(BigInteger.valueOf(Long.parseLong(firstStone)))) {
//                res1 = memo.get(BigInteger.valueOf(Long.parseLong(firstStone)));
//            } else {
//                res1 = blink_part_2(BigInteger.valueOf(Long.parseLong(firstStone)), memo, score, blinks - 1);
//            }
            BigInteger res1;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(firstStone)), blinks - 1))) {
                res1 = memo.get(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(firstStone)), blinks - 1));
            } else {
                res1 = blink_part_2(BigInteger.valueOf(Long.parseLong(firstStone)), memo, score, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(firstStone)), blinks - 1), res1);
            }

            //var res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(secondStone)), values, memo, score, blinks - 1);

//            BigInteger res2;
//            if (memo.containsKey(BigInteger.valueOf(Long.parseLong(secondStone)))) {
//                res2 = memo.get(BigInteger.valueOf(Long.parseLong(secondStone)));
//            } else {
//                res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(secondStone)), memo, score, blinks - 1);
//            }
            BigInteger res2;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(secondStone)), blinks - 1))) {
                res2 = memo.get(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(secondStone)), blinks - 1));
            } else {
                res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(secondStone)), memo, score, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(secondStone)), blinks - 1), res2);
            }

//            if (!memo.containsKey(BigInteger.valueOf(Long.parseLong(firstStone)))) {
//                values.put(firstStone, res1);
//            }
//
//            if (!memo.containsKey(BigInteger.valueOf(Long.parseLong(secondStone)))) {
//                values.put(secondStone, res2);
//            }
            score = score.add(res1);
            score = score.add(res2);
            //System.out.println("Rule two - Number: " + number + " blink: " + blinks+ " score: " + score);
//            if (!this.memo.containsKey(BigInteger.valueOf(Long.parseLong(stoneString))) && !res1.equals(BigInteger.ZERO)) {
//                this.memo.put(BigInteger.valueOf(Long.parseLong(stoneString)), score);
//            }
        } else {
//            BigInteger res2;
//            if (memo.containsKey(number)) {
//                res2 = memo.get(BigInteger.valueOf(Long.parseLong(stoneString) * 2024));
//            } else {
//                res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), memo, score, blinks - 1);
//            }
            BigInteger res2;
            if (memo.containsKey(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), blinks - 1))) {
                res2 = memo.get(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), blinks - 1));
            } else {
                res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), memo, score, blinks - 1);
                memo.put(new AbstractMap.SimpleEntry<>(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), blinks - 1), res2);
            }
            //var res2 = blink_part_2(BigInteger.valueOf(Long.parseLong(stoneString) * 2024), values, memo, score, blinks - 1);
//            if (!memo.containsKey(BigInteger.valueOf(Long.parseLong(stoneString) * 2024))) {
//                values.put(BigInteger.valueOf(Long.parseLong(stoneString) * 2024).toString(), res2);
//            }
            score = score.add(res2);
            //System.out.println("Rule third - Number: " + number + " blink: " + blinks + " score: " + score);
        }

//        if (!memo.containsKey(BigInteger.valueOf(Long.parseLong(stoneString)))) {
//            values.put(BigInteger.valueOf(Long.parseLong(stoneString)).toString(), score);
//        }
        //memo.put(number, score);
        //System.out.println(memo);

        return score;

    }

    public int Part1(int times, ArrayList<BigInteger> stones) {

        for(int i = 0; i < times; i++) {
            stones = blink_part_1(stones);
            System.out.println(stones);
        }

        return stones.size();
    }

    public BigInteger Part2(int times, ArrayList<BigInteger> stones) {

        // Initialising the memo
        BigInteger result = BigInteger.ZERO;

        System.out.println(this.memo);

        HashMap<AbstractMap.SimpleEntry<BigInteger, Integer>, BigInteger> memo = new HashMap<>();

        for(BigInteger stone : stones) {
            var temp = blink_part_2(stone, memo, BigInteger.ZERO, times);
            result = result.add(temp);
        }

        return result;
    }


}
