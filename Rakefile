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

desc "mk leetcode playground as: num=7 title=reverse-integer func=reverse"
task :leetcode do
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
  if File.directory?(target_folder) && File.exist?(target_folder)
    abort("rake aborted! #{target_folder} not overwrite") if ask("#{target_folder} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
  else
      mkdir_p target_folder
  end
  target_src_file = File.join(target_folder, "#{now_mark}.#{CONFIG['go_ext']}")
  if File.exist?(target_src_file)
    abort("rake aborted! #{target_src_file} not overwrite") if ask("#{target_src_file} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
  end
  puts "Creating new source: #{target_src_file}"
  open(target_src_file, 'w') do |post|
    post.puts "package #{target_package_name}"
    post.puts ""
    post.puts "func #{problems_func}(s string) int {"
    post.puts "}"
  end

  target_test_file = File.join(target_folder, "#{now_mark}_test.#{CONFIG['go_ext']}")
  if File.exist?(target_test_file)
    abort("rake aborted! #{target_test_file} not overwrite") if ask("#{target_test_file} already exists. Do you want to overwrite?", ['y', 'n']) == 'n'
  end
  puts "Creating new unit test: #{target_test_file}"
  open(target_test_file, 'w') do |post|
    post.puts "package #{target_package_name}"
    post.puts ""
    post.puts "import ("
    post.puts "	\"github.com/stretchr/testify/assert\""
    post.puts "	\"testing\""
    post.puts ")"
    post.puts ""
    post.puts "func Test_#{problems_func}(t *testing.T) {"
    post.puts "	// mock #{problems_func}"
    post.puts "	qs := []question{"
    post.puts "		{"
    post.puts "			param{\"III\"},"
    post.puts "			answer{3},"
    post.puts "		},"
    post.puts "	}"
    post.puts "	t.Logf(\"~> LeetCode #{problems_func} start\")"
    post.puts "	// do #{problems_func}"
    post.puts "	for _, q := range qs {"
    post.puts "		a, p := q.answer, q.param"
    post.puts "		// verify #{problems_func}"
    post.puts "		assert.Equal(t, a.one, #{problems_func}(p.one),"
    post.puts "			\"in [ %v ], out [%v]\", p.one, #{problems_func}(p.one))"
    post.puts "	}"
    post.puts "	t.Logf(\"~> LeetCode #{problems_func} end\")"
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