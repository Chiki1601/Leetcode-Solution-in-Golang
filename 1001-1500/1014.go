func maxScoreSightseeingPair(values []int) int {
    max := 0;
    score := values[0];
    for i := 1; i < len(values); i++ {
      score--;
      if (score + values[i] > max) {
        max = score + values[i];
      }
      if (values[i] > score) {
        score = values[i];
      }
    }
    return max;
}
