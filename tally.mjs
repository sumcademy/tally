// How to run:
// node tally.mjs
const COUNT = 50_000_000;

// fill `positives` array with numbers from 1 to 50,000,000
const positives = [];
for (let i = 1; i <= COUNT; i++) {
  positives.push(i);
}

// fill `negatives` array with numbers from -1 to -50,000,000
const negatives = [];
for (let i = -1; i >= -COUNT; i--) {
  negatives.push(i);
}

let total = 0;
// adds all numbers in `numbers` into `total`
async function tally(numbers) {
  console.log("  inner job: tallying", numbers.length, "numbers...");
  const startTime = Date.now();
  for (const n of numbers) { total += n }
  const totalTime = Date.now() - startTime;
  console.log("    took", totalTime / 1000, "seconds");
}


console.log("outer job: tallying positives and negatives...");
const startTime = Date.now();
await Promise.all([tally(positives), tally(negatives)]);
const totalTime = Date.now() - startTime;
console.log("outer job took", totalTime / 1000, "seconds total");

console.log("total sum:", total);
