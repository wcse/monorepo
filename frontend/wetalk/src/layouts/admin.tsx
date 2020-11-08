import Footer from 'components/layout/Footer';
import Header from 'components/layout/Header';

const AdminLayout = ({ children }) => {
  return (
    <div>
      <Header></Header>
      <div className="container mx-auto">{children}</div>
      <Footer></Footer>
    </div>
  );
};

export default AdminLayout;
