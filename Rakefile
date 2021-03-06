require 'rake'

SOURCE = "."

CONFIG = {
  'leetcode_playground_path' => File.join(SOURCE, "playground"),
  'go_ext' => "go",
}

def ask(message, valid_options)
  if valid_options
    answer = get_stdin("#{message} #{valid_options.to_s.gsub(/"/, '').gsub(/, /,'/')} ") while !valid_options.include?(answer)
  else
    answer = get_stdin(message)
  end
  answer
end

def get_stdin(message)
  print message
  STDIN.gets.chomp
end

namespace :leetcode do

  desc "mk leetcode playground folder env: num=7 title=reverse-integer"
  task :mkdir do
    abort("rake aborted: '#{CONFIG['leetcode_playground_path']}' directory not found.") unless FileTest.directory?(CONFIG['leetcode_playground_path'])
    problems_num = ENV["num"] || "0"
    problems_title = ENV["title"] || "two-sum"
    if problems_num == "0"
      abort("rake aborted: must set num=1, now is: #{problems_num}")
    end
    if problems_title == ""
      abort("rake aborted: must set title="", now is: #{problems_title}")
    end
    problems_num = problems_num.rjust(5, "0")
    problems_title = problems_title.downcase.strip.gsub(' ', '-')
    now_mark = "#{problems_num}-#{problems_title}"
    target_package_name = "_#{now_mark.gsub('-', '_')}"
    target_folder = File.join(CONFIG['leetcode_playground_path'], now_mark)
    if not File.exist?(target_folder)
      mkdir_p target_folder
    end
  end

  desc "mk leetcode playground source env: num=7 title=reverse-integer func=reverse"
  task :src do
    abort("rake aborted: '#{CONFIG['leetcode_playground_path']}' directory not found.") unless FileTest.directory?(CONFIG['leetcode_playground_path'])
    problems_num = ENV["num"] || "0"
    problems_title = ENV["title"] || "two-sum"
    problems_func = ENV["func"] || "twoSum"
    if problems_num == "0"
      abort("rake aborted: must set num=1, now is: #{problems_num}")
    end
    if problems_title == ""
      abort("rake aborted: must set title="", now is: #{problems_title}")
    end
    if problems_func == ""
      abort("rake aborted: must set func="", now is: #{problems_func}")
    end
    problems_num = problems_num.rjust(5, "0")
    problems_title = problems_title.downcase.strip.gsub(' ', '-')
    now_mark = "#{problems_num}-#{problems_title}"
    target_package_name = "_#{now_mark.gsub('-', '_')}"
    target_folder = File.join(CONFIG['leetcode_playground_path'], now_mark)
    if not File.exist?(target_folder)
      mkdir_p target_folder
    end
    target_src_file = File.join(target_folder, "#{now_mark}.#{CONFIG['go_ext']}")
    if File.exist?(target_src_file)
      abort("rake aborted! #{target_src_file} not overwrite") if ask("#{target_src_file} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
    end
    puts "Creating #{problems_title} source: #{target_src_file}"
    open(target_src_file, 'w:utf-8') do |post|
      post.puts "package #{target_package_name}"
      post.puts ""
      post.puts "func #{problems_func}(s string) int {"
      post.puts "}"
    end
  end

  desc "mk leetcode playground unit test env: num=7 title=reverse-integer func=reverse"
  task :test do
    abort("rake aborted: '#{CONFIG['leetcode_playground_path']}' directory not found.") unless FileTest.directory?(CONFIG['leetcode_playground_path'])
    problems_num = ENV["num"] || "0"
    problems_title = ENV["title"] || "two-sum"
    problems_func = ENV["func"] || "twoSum"
    if problems_num == "0"
      abort("rake aborted: must set num=1, now is: #{problems_num}")
    end
    if problems_title == ""
      abort("rake aborted: must set title="", now is: #{problems_title}")
    end
    if problems_func == ""
      abort("rake aborted: must set func="", now is: #{problems_func}")
    end
    problems_num = problems_num.rjust(5, "0")
    problems_title = problems_title.downcase.strip.gsub(' ', '-')
    now_mark = "#{problems_num}-#{problems_title}"
    target_package_name = "_#{now_mark.gsub('-', '_')}"
    target_folder = File.join(CONFIG['leetcode_playground_path'], now_mark)
    if not File.exist?(target_folder)
      mkdir_p target_folder
    end
    target_test_file = File.join(target_folder, "#{now_mark}_test.#{CONFIG['go_ext']}")
    if File.exist?(target_test_file)
      abort("rake aborted! #{target_test_file} not overwrite") if ask("#{target_test_file} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
    end
    puts "Creating #{problems_title} unit test: #{target_test_file}"
    open(target_test_file, 'w:utf-8') do |post|
      post.puts "package #{target_package_name}"
      post.puts ""
      post.puts "import ("
      post.puts "	\"github.com/stretchr/testify/assert\""
      post.puts "	\"testing\""
      post.puts ")"
      post.puts ""
      post.puts "	// mock #{problems_func}"
      post.puts ""
      post.puts "var qs = []question{"
      post.puts "	{"
      post.puts "		param{\"III\"},"
      post.puts "		answer{3},"
      post.puts "	},"
      post.puts "}"
      post.puts ""
      post.puts "func Test_#{problems_func}(t *testing.T) {"
      post.puts "	t.Logf(\"~> LeetCode #{problems_func} start\")"
      post.puts "	// do #{problems_func}"
      post.puts "	for _, q := range qs {"
      post.puts "		a, p := q.answer, q.param"
      post.puts "		res := #{problems_func}(p.one)"
      post.puts "		// verify #{problems_func}"
      post.puts "		assert.Equal(t, a.one, res,"
      post.puts "			\"fail: in [ %v ], out [%v] , want [ %v ]\", p.one, res, a.one)"
      post.puts "	}"
      post.puts "	t.Logf(\"~> LeetCode #{problems_func} end\")"
      post.puts "}"
      post.puts ""
      post.puts "func Benchmark_#{problems_func}(b *testing.B) {"
      post.puts "	// do #{problems_func}"
      post.puts "	for _, q := range qs {"
      post.puts "		#{problems_func}(q.param.one)"
      post.puts "	}"
      post.puts "}"
      post.puts ""
      post.puts "type question struct {"
      post.puts "	param"
      post.puts "	answer"
      post.puts "}"
      post.puts ""
      post.puts "// one first param"
      post.puts "type param struct {"
      post.puts "	one string"
      post.puts "}"
      post.puts ""
      post.puts "// one first  answer"
      post.puts "type answer struct {"
      post.puts "	one int"
      post.puts "}"
    end
  end

  desc "mk leetcode playground answer document env: num=7 title=reverse-integer func=reverse"
  task :answer do
    abort("rake aborted: '#{CONFIG['leetcode_playground_path']}' directory not found.") unless FileTest.directory?(CONFIG['leetcode_playground_path'])
    problems_num = ENV["num"] || "0"
    problems_title = ENV["title"] || "two-sum"
    problems_func = ENV["func"] || "twoSum"
    if problems_num == "0"
      abort("rake aborted: must set num=1, now is: #{problems_num}")
    end
    if problems_title == ""
      abort("rake aborted: must set title="", now is: #{problems_title}")
    end
    if problems_func == ""
      abort("rake aborted: must set func="", now is: #{problems_func}")
    end
    problems_num = problems_num.rjust(5, "0")
    problems_title = problems_title.downcase.strip.gsub(' ', '-')
    now_mark = "#{problems_num}-#{problems_title}"
    target_package_name = "_#{now_mark.gsub('-', '_')}"
    target_folder = File.join(CONFIG['leetcode_playground_path'], now_mark)
    if not File.exist?(target_folder)
      mkdir_p target_folder
    end
    target_answer_file = File.join(target_folder, "README.md")
    puts "Creating #{problems_title} answer documentation: #{target_answer_file}"
    if File.exist?(target_answer_file)
      abort("rake aborted! #{target_answer_file} not overwrite") if ask("#{target_answer_file} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
    end
    open(target_answer_file, 'w:utf-8') do |post|
      post.puts "### 解题思路"
      post.puts ""
      post.puts ""
      post.puts "### 代码"
      post.puts ""
      post.puts "```go"
      post.puts "func #{problems_func}() {"
      post.puts "}"
      post.puts "```"
    end
  end
end

desc "playground env: num=7 title=reverse-integer func=reverse"
task :leetcode => ['leetcode:src', 'leetcode:answer', 'leetcode:test']