require 'rails_helper'

describe Price::Calculate do
  subject { described_class.run!(params) }

  let(:tariff) { nil }
  let(:params) do
    {
      start_point: {
        lat: 55.729060,
        lng: 37.622691
      },
      end_point: {
        lat: 55.808116,
        lng: 37.581609
      },
      tariff: tariff
    }
  end

  before do
    result = {
      distance: 12135.455,
      time: 1284809
    }

    allow(Direction::Api).to receive(:calculate).and_return(result)
  end

  context 'when default tariff' do
    it 'returns price' do
      expect(subject).to eq(4775.to_d)
    end
  end

  context 'when custom tariff' do
    let(:tariff) do
      {
        min_price: 500.0,
        order_price: 250.0,
        minute_price: 20.0,
        km_price: 30.0
      }
    end

    it 'returns price' do
      expect(subject).to eq(4897.to_d)
    end
  end
end
