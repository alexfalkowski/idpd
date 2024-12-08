# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'idpd/http'

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

    def http
      @http ||= Idpd::HTTP.new('http://localhost:11000')
    end
  end
end
