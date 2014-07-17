require 'guard/plugin'

module ::Guard
  class Go < ::Guard::Plugin
    def run_all
      system "go test"
    end

    def run_on_change(paths)
      run_all 
    end
  end
end

guard :go do
  watch(%r(.*.go))
end
