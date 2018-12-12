require 'rails_helper'

describe Api::V1::PricesController, type: :request do
  let(:params) { {} }
  let(:headers) do
    {
      'Content-Type' => 'application/json',
    }
  end

  before do
    result = {
      distance: 12135.455,
      time: 1284809
    }

    allow(Direction::Api).to receive(:calculate).and_return(result)
  end

  context 'POST #calculate' do
    subject do
      post '/api/v1/prices/calculate', params: params.to_json, headers: headers
    end

    let(:params) do
      {
        start_point: {
          lat: 55.729060,
          lng: 37.622691
        },
        end_point: {
          lat: 55.808116,
          lng: 37.581609
        }
      }
    end

    it 'returns ok' do
      subject
      expect(response).to have_http_status(:success)
    end

    it 'returns json' do
      subject
      expect(response.body).to eq({ price: 4775.0 }.to_json)
    end
  end
end
