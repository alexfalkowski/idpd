# frozen_string_literal: true

Given('we create an invalid pipeline') do
  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => ['hellos "1"', 'echo "2"'] }] }
  }

  @response = Idpd.http.create_pipeline(pipeline.to_json, opts)
end

When('we create a simple pipeline') do
  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => ['echo "1"', 'echo "2"'] }] }
  }

  @response = Idpd.http.create_pipeline(pipeline.to_json, opts)
end

When('we try to create a invalid pipeline of kind {string}') do |kind|
  cases = {
    'invalid pipeline name' => { pipeline: {} },
    'missing jobs' => { pipeline: { 'name' => 'test' } },
    'invalid job name' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'steps' => %w[test test2] }] } },
    'missing steps' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test' }] } }
  }

  @response = Idpd.http.create_pipeline(cases[kind].to_json, opts)
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

  @response = Idpd.http.create_pipeline(pipeline.to_json, opts)
end

When('we try to create a pipeline with a bad payload') do
  @response = Idpd.http.create_pipeline(Base64.encode64('test'), opts)
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
