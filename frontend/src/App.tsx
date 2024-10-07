import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import ContractList from './pages/ContractList';
import BillDetail from './pages/BillDetail';

const App: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Dashboard} />
        <Route path="/contracts" component={ContractList} />
        <Route path="/bills/:id" component={BillDetail} />
      </Switch>
    </Router>
  );
};

export default App;