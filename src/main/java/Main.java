import java.sql.*;


public class Main {


    public void search(String id) throws Exception {


        Connection conn = DriverManager.getConnection(
                "jdbc:mysql://localhost/test",
                "root",
                "password"
        );


        String sql =
            "select * from users where id=" + id;


        Statement stmt = conn.createStatement();


        stmt.executeQuery(sql);

    }


}