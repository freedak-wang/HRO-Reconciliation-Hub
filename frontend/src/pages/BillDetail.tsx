import React from 'react';
import { useParams } from 'react-router-dom';
import { Typography, Descriptions, Table } from 'antd';

const { Title } = Typography;

const BillDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  // TODO: Fetch bill details

  return (
    <div>
      <Title level={2}>Bill Details</Title>
      <Descriptions title="Bill Information">
        <Descriptions.Item label="Bill ID">{id}</Descriptions.Item>
        {/* Add more bill information */}
      </Descriptions>
      {/* TODO: Add tables for personnel services and non-personal items */}
      {/* TODO: Add section for differences and adjustments */}
    </div>
  );
};

export default BillDetail;