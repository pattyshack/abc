load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

abseil_cpp_commit_hash = "40070892631b82c79864ff4d3f7e12105f003157"
googletest_commit_hash = "be03d00f5f0cc3a997d1a368bee8a1fe93651f48"
benchmark_commit_hash = "aacf2b1af967b083565be8c7181626b4609318ac"

http_archive(
  name = "bazel_skylib",
  urls = ["https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz"],
  sha256 = "66ffd9315665bfaafc96b52278f57c7e2dd09f5ede279ea6d39b2be471e7e3aa",
)

http_archive(
  name = "com_google_absl",
  urls = ["https://github.com/abseil/abseil-cpp/archive/" + abseil_cpp_commit_hash + ".zip"],
  strip_prefix = "abseil-cpp-" + abseil_cpp_commit_hash,
)

http_archive(
  name = "com_google_googletest",
  urls = ["https://github.com/google/googletest/archive/" + googletest_commit_hash + ".zip"],
  strip_prefix = "googletest-" + googletest_commit_hash,
)

http_archive(
  name = "com_github_google_benchmark",
  urls = ["https://github.com/google/benchmark/archive/" + benchmark_commit_hash + ".zip"],
  strip_prefix = "benchmark-" + benchmark_commit_hash,
)

