# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'idpd/v1/http'

module Idpd
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:11000')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.yml')
    end

    def token
      Nonnative::Header.auth_bearer(Base64.decode64(File.read('secrets/token')))
    end
  end

  module V1
    class << self
      def server
        @server ||= Idpd::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
