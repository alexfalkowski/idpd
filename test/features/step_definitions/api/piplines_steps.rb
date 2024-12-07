# frozen_string_literal: true

When('we create a simple pipeline') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }

  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.create_pipeline(pipeline, opts)
end

When('we try to create a invalid pipeline of kind {string}') do |kind|
  cases = {
    'invalid pipeline name' => { pipeline: {} },
    'missing jobs' => { pipeline: { 'name' => 'test' } },
    'invalid job name' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'steps' => %w[test test2] }] } },
    'missing steps' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test' }] } }
  }

  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }

  @response = Idpd::V1.server.create_pipeline(cases[kind], opts)
end

When('we try to create a pipeline with being authorized') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.create_pipeline(pipeline, opts)
end

When('we get the created simple pipeline') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }

  @response = Idpd::V1.server.get_pipeline(2, opts)
end

When('we get an nonexistent pipeline') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }

  @response = Idpd::V1.server.get_pipeline(10, opts)
end

When('we get an invalid pipeline') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }

  @response = Idpd::V1.server.get_pipeline(0, opts)
end

Then('we should have a created pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(1)
end

Then('we should have bad request') do
  expect(@response.code).to eq(400)
end

Then('we should have a unauthorized request') do
  expect(@response.code).to eq(401)
end

Then('we should have a simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(2)
end

Then('we should have a not found request') do
  expect(@response.code).to eq(404)
end
