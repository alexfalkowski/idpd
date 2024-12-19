# frozen_string_literal: true

When('we trigger the simple pipeline') do
  @id = JSON.parse(@response.body)['pipeline']['id']
  @response = Idpd.http.trigger_pipeline(@id, Idpd.options)
end

When('we try to trigger an nonexistent pipeline') do
  @response = Idpd.http.trigger_pipeline('123456', Idpd.options)
end

When('we try to trigger with an invalid id') do
  @response = Idpd.http.trigger_pipeline(0, Idpd.options)
end

Then('we should receive the output of simple pipeline') do
  expect(@response.code).to eq(200)

  steps = JSON.parse(@response.body)['pipeline']['jobs'][0]['steps']
  expect(steps).to eq(['"1"', '"2"'])
end

Then('we should receive that the pipeline contained errors') do
  expect(@response.code).to eq(500)
  expect(@response.body.strip).to eq('rest: pipeline test: job test failed: step hellos "1": exec: "hellos": executable file not found in $PATH')
end
