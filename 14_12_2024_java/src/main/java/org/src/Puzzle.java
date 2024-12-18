package org.src;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

import java.io.*;
import java.nio.charset.Charset;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Getter
@Setter
public class Puzzle {

    // Used path
    private String path;

    // Robots
    private ArrayList<Robot> robots;

    // Bounds
    private int rows;
    private int cols;

    private int excludedY;
    private int excludedX;

    // Time to run the sim
    private int seconds;
    private int highest;

    @Getter
    @Setter
    @AllArgsConstructor
    public class Robot {
        private int posX;
        private int posY;

        private int velocityX;
        private int velocityY;

        // Walking and ...
        public AbstractMap.SimpleEntry<Integer, Integer> walk() {
            this.posX = Math.floorMod((this.posX + this.velocityX), (cols));
            this.posY = Math.floorMod((this.posY + this.velocityY), (rows));

            // For testing purposes
            return new AbstractMap.SimpleEntry<>(posY, posX);
        }

        public AbstractMap.SimpleEntry<Integer, Integer> coordsAfter(int time) {
//            this.posX = Math.floorMod((this.posX + this.velocityX * time), (cols));
//            this.posY = Math.floorMod((this.posY + this.velocityY * time), (rows));
            return new AbstractMap.SimpleEntry<>(Math.floorMod((this.posY + this.velocityY * time), (rows)),
                    Math.floorMod((this.posX + this.velocityX * time), (cols)));
        }
    }

    public Puzzle(String filePath, int rows, int cols, int seconds) {
        this.path = filePath;
        this.highest = 0;

        // Bounds
        this.rows = rows;
        this.cols = cols;
        this.excludedY = (rows - 1) / 2;
        this.excludedX = (cols - 1) / 2;


        // Repetition
        this.seconds = seconds;

        // Data
        this.robots = new ArrayList<>();

        // Reading the input
        try (var in = Puzzle.class.getResourceAsStream(this.path)) {

            String data = new String(in.readAllBytes(), Charset.defaultCharset());
            Scanner myReader = new Scanner(data);

            Pattern pattern = Pattern.compile("[-]*\\d+");
            // Reading the data and calculating from the get go
            while (myReader.hasNextLine()) {
                // Total of 4 integers to read, the pos and the vel
                Matcher intMatcher = pattern.matcher(myReader.nextLine());
                ArrayList<Integer> foundValues = new ArrayList<>();

                // Searching the line for the aformentioned four value
                while(intMatcher.find()) {
                    foundValues.add(Integer.valueOf(intMatcher.group()));
                }

                this.robots.add(new Robot(foundValues.get(0),
                        foundValues.get(1), foundValues.get(2), foundValues.get(3)));
            }

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public int Part1() {

        int ul_quadrant = 0;
        int ur_quadrant = 0;
        int ll_quadrant = 0;
        int lr_quadrant = 0;

        for(Robot r : this.robots) {
            var coords = r.coordsAfter(this.seconds);
            //System.out.println("Point (" + coords.getKey() + "," + coords.getValue() + "), limit y " + this.excludedX + " limit x" + this.excludedY);
            // Upper left quadrant
            if (coords.getKey() < this.excludedY && coords.getValue() < this.excludedX) {
                ul_quadrant++;
            } else if (coords.getKey() < this.excludedY && coords.getValue() > this.excludedX) {
                ur_quadrant++;
            } else if (coords.getKey() > this.excludedY && coords.getValue() < this.excludedX) {
                ll_quadrant++;
            } else if (coords.getKey() > this.excludedY && coords.getValue() > this.excludedX) {
                lr_quadrant++;
            }
        }
        //System.out.println("Size: " + this.robots.size());
        //System.out.println(ul_quadrant + " " + ur_quadrant + " " + ll_quadrant + " " + lr_quadrant);
        return ul_quadrant * ur_quadrant * ll_quadrant * lr_quadrant;
    }

    public Set<AbstractMap.SimpleEntry<Integer, Integer>> makeChristmasTreeCoordinates() {
        Set<AbstractMap.SimpleEntry<Integer, Integer>> coordinates = new HashSet<>();

        // Creating the 'tree'
        var y = 0;
        var middleX = excludedX;

        coordinates.add(new AbstractMap.SimpleEntry<>(y, excludedX));
        coordinates.add(new AbstractMap.SimpleEntry<>(++y, excludedX));

        var counter = 1;
        while(y + counter < this.excludedY || middleX + counter < this.cols || middleX - counter > 0) {
            coordinates.add(new AbstractMap.SimpleEntry<>(y + counter, middleX - counter));
            coordinates.add(new AbstractMap.SimpleEntry<>(y + counter, middleX + counter));
            counter++;
        }

        // Creating the bottom of the tree
        y = excludedY;
        counter = 1;

        while(middleX - counter > 0 || middleX + counter < this.cols) {
            coordinates.add(new AbstractMap.SimpleEntry<>(y, middleX - counter));
            coordinates.add(new AbstractMap.SimpleEntry<>(y, middleX + counter));
            counter++;
        }

        // Creating the bottom of the tree
        y = excludedY;
        counter = 1;

        while(y + counter < this.rows) {
            coordinates.add(new AbstractMap.SimpleEntry<>(y + counter, middleX - 1));
            coordinates.add(new AbstractMap.SimpleEntry<>(y + counter, middleX + 1));
            counter++;
        }

        return coordinates;
    }

    public boolean checkForChrismasTree(Set<AbstractMap.SimpleEntry<Integer, Integer>> coords, int sec) {
//        Set<AbstractMap.SimpleEntry<Integer, Integer>> points = new HashSet<>();
//        for(Robot r: this.robots) {
//            points.add(new AbstractMap.SimpleEntry<>(r.getPosY(), r.getPosX()));
//        }
//
//
//        this.highest = continuous;


//        //Not working, but the quickest
//        String[][] tree = new String[rows][cols];
//        HashSet<AbstractMap.SimpleEntry<Integer, Integer>> points = new HashSet<>();
//
//        for (Robot rn : this.robots) {
//            tree[rn.getPosY()][rn.getPosX()] = "X";
//            points.add(new AbstractMap.SimpleEntry<>(rn.getPosY(), rn.getPosX()));
//        }
//
//        return points.size() == this.robots.size();


//        //Checking for vertical line
//        boolean result = false;
//        int continuous = 0;
//        for (Robot r : this.robots) {
//            for (Robot j : this.robots) {
//                if (r.getPosY() == j.getPosY() && Math.abs(r.getPosX() - j.getPosX()) == 1) {
//                    continuous++;
//                }
//            }
//        }
//        //System.out.println(continuous);
//        if (continuous >= this.highest - 10) {
//            this.highest = continuous;
//
//            String[][] tree = new String[this.rows][this.cols];
//
//            for (Robot rn : this.robots) {
//                tree[rn.getPosY()][rn.getPosX()] = "X";
//            }
//
//            System.out.println("SIEKONDS: " + sec);
//            for (int ik = 0; ik < this.rows; ik++) {
//                for (int jk = 0; jk < this.cols; jk++) {
//                    if (!Objects.equals(tree[ik][jk], "X")) {
//                        tree[ik][jk] = ".";
//                    }
//                    System.out.print(tree[ik][jk] + " ");
//                }
//                System.out.print("\n");
//            }
//        }
//
        return false;
    }

    public int Part2() throws InterruptedException {

//        int res = 0;
//
////        var coordinates = makeChristmasTreeCoordinates();
////        System.out.println(coordinates);
//
//        // For drawing the tree
//        //String[][] tree = new String[rows][cols];
//
////        for(AbstractMap.SimpleEntry<Integer, Integer> e: coordinates) {
////            tree[e.getKey()][e.getValue()] = 1;
////        }
//
////        for(int i = 0; i < rows; i++) {
////            for( int j = 0; j < cols; j++) {
////                if (tree[i][j] != "X") {
////                    tree[i][j] = ".";
////                }
////                System.out.print(tree[i][j] + " ");
////            }
////            System.out.print("\n");
////        }
//
//        while(!checkForChrismasTree(null, res)) {
//            //String[][] tree = new String[rows][cols];
//
////            for(Robot r : this.robots) {
////                tree[r.getPosY()][r.getPosX()] = "X";
////            }
////
////            for(int i = 0; i < rows; i++) {
////                for( int j = 0; j < cols; j++) {
////                    if (!Objects.equals(tree[i][j], "X")) {
////                        tree[i][j] = ".";
////                    }
////                    System.out.print(tree[i][j] + " ");
////                }
////                System.out.print("\n");
////            }
//            //Thread.sleep(500);
//
//            for(Robot r: this.robots) {
//                r.walk();
//            }
//            res++;
//        }
//
//        return res;
//
//        for(int i = 0; i <= 7860; i++) {
//            this.robots.forEach(Robot::walk);
//        }
        //this.robots.forEach(r -> r.coordsAfter(78411));

//        String[][] tree = new String[rows][cols];
////
//            for(Robot r : this.robots) {
//                tree[r.getPosY()][r.getPosX()] = "X";
//            }

            int score = Integer.MAX_VALUE;
            int when = 0;
            for(int x = 1; x < this.rows * this.cols; x++) {
                this.seconds = x;

                int safety = Part1();
                if (safety < score) {
                    score = safety;
                    when = x;
                }
            }
//
//                for (int i = 1; i < rows; i++) {
//                    for (int j = 1; j < cols; j++) {
//                        if (!Objects.equals(tree[i][j], "X")) {
//                            tree[i][j] = ".";
//                        }
//                        System.out.print(tree[i][j] + " ");
//                    }
//                    System.out.print("\n");
//                }


//        this.seconds = 1;
//        int when = 0;
//        int score = Integer.MAX_VALUE;
//        for(int i = 0; i < this.rows * this.cols; i++) {
//            int safety = Part1();
//            if(safety < score) {
//                score = safety;
//                when = i;
//            }
//        }
//
//        return when;
        return when;
    }

}
