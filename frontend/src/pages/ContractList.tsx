import React from 'react';
import { Table, Typography } from 'antd';

const { Title } = Typography;

const ContractList: React.FC = () => {
  // TODO: Fetch contracts data

  const columns = [
    {
      title: 'Contract ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: 'Client',
      dataIndex: 'client',
      key: 'client',
    },
    {
      title: 'Start Date',
      dataIndex: 'startDate',
      key: 'startDate',
    },
    {
      title: 'End Date',
      dataIndex: 'endDate',
      key: 'endDate',
    },
    // Add more columns as needed
  ];

  return (
    <div>
      <Title level={2}>Contracts</Title>
      <Table columns={columns} /* dataSource={contracts} */ />
    </div>
  );
};

export default ContractList;