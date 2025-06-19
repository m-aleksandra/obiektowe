import unittest
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options


class UITests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        options = Options()
        # options.add_argument("--headless")
        cls.driver = webdriver.Chrome(options=options)
        cls.driver.get("http://localhost:5000/")
        cls.driver.implicitly_wait(3)

    @classmethod
    def tearDownClass(cls):
        cls.driver.quit()

    def setUp(self):
        self.driver.get("http://localhost:5000/")

    def test_title(self):
        self.assertIn("Welcome", self.driver.title)

    def test_input_visible(self):
        input_el = self.driver.find_element(By.ID, "name")
        self.assertTrue(input_el.is_displayed())

    def test_button_visible(self):
        button_el = self.driver.find_element(By.TAG_NAME, "button")
        self.assertTrue(button_el.is_displayed())

    def test_output_empty_on_load(self):
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output.strip(), "")

    def test_hello_john(self):
        self.driver.find_element(By.ID, "name").send_keys("John")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.5)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, John!")

    def test_hello_empty(self):
        self.driver.find_element(By.ID, "name").clear()
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.5)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, stranger!")


    def test_user_page(self):
        self.driver.get("http://localhost:5000/user/Marta")
        self.assertIn("User page for: Marta", self.driver.page_source)

    def test_placeholder_text(self):
        placeholder = self.driver.find_element(By.ID, "name").get_attribute("placeholder")
        self.assertEqual(placeholder, "Enter your name")

    def test_button_text(self):
        button = self.driver.find_element(By.TAG_NAME, "button")
        self.assertEqual(button.text, "Say Hello")

    def test_input_has_correct_type(self):
        input_el = self.driver.find_element(By.ID, "name")
        self.assertEqual(input_el.get_attribute("type"), "text")

    def test_input_has_id_name(self):
        input_el = self.driver.find_element(By.ID, "name")
        self.assertEqual(input_el.get_attribute("id"), "name")

    def test_output_has_correct_tag(self):
        output_el = self.driver.find_element(By.ID, "output")
        self.assertEqual(output_el.tag_name.lower(), "p")

    def test_button_tag_is_button(self):
        button_el = self.driver.find_element(By.TAG_NAME, "button")
        self.assertEqual(button_el.tag_name.lower(), "button")

    def test_header_text(self):
        h1_el = self.driver.find_element(By.TAG_NAME, "h1")
        self.assertEqual(h1_el.text.strip(), "Welcome!")

    def test_input_placeholder_not_empty(self):
        input_el = self.driver.find_element(By.ID, "name")
        self.assertTrue(input_el.get_attribute("placeholder"))

    def test_name_input_clearable(self):
        input_el = self.driver.find_element(By.ID, "name")
        input_el.send_keys("something")
        input_el.clear()
        self.assertEqual(input_el.get_attribute("value"), "")

    def test_button_click_does_not_crash(self):
        button_el = self.driver.find_element(By.TAG_NAME, "button")
        button_el.click()
        output = self.driver.find_element(By.ID, "output").text
        self.assertTrue(output.startswith("Hello"))

    def test_user_page_direct_2(self):
        self.driver.get("http://localhost:5000/user/Ala")
        self.assertIn("User page for: Ala", self.driver.page_source)

    def test_user_page_direct_3(self):
        self.driver.get("http://localhost:5000/user/Bob")
        self.assertIn("User page for: Bob", self.driver.page_source)

    def test_user_page_case_sensitive(self):
        self.driver.get("http://localhost:5000/user/alice")
        self.assertIn("User page for: alice", self.driver.page_source)

    def test_hello_uppercase(self):
        self.driver.find_element(By.ID, "name").send_keys("TOM")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.2)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, TOM!")

    def test_hello_with_spaces(self):
        self.driver.find_element(By.ID, "name").send_keys("   Jan  ")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.2)
        output = self.driver.find_element(By.ID, "output").text
        self.assertIn("Jan", output)

    def test_input_accepts_digits(self):
        self.driver.find_element(By.ID, "name").send_keys("12345")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.2)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, 12345!")

    def test_input_accepts_symbols(self):
        self.driver.find_element(By.ID, "name").send_keys("@@@")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.2)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, @@@!")

    def test_input_accepts_unicode(self):
        self.driver.find_element(By.ID, "name").send_keys("Łukasz")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.2)
        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, Łukasz!")


    def test_output_is_updated(self):
        self.driver.find_element(By.ID, "name").send_keys("One")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.5)

        self.driver.find_element(By.ID, "name").clear()
        self.driver.find_element(By.ID, "name").send_keys("Two")
        self.driver.find_element(By.TAG_NAME, "button").click()
        time.sleep(0.5) 

        output = self.driver.find_element(By.ID, "output").text
        self.assertEqual(output, "Hello, Two!")


    def test_button_exists_once(self):
        buttons = self.driver.find_elements(By.TAG_NAME, "button")
        self.assertEqual(len(buttons), 1)

    def test_input_exists_once(self):
        inputs = self.driver.find_elements(By.ID, "name")
        self.assertEqual(len(inputs), 1)

    def test_output_exists_once(self):
        outputs = self.driver.find_elements(By.ID, "output")
        self.assertEqual(len(outputs), 1)

    def test_script_is_present(self):
        scripts = self.driver.find_elements(By.TAG_NAME, "script")
        self.assertGreaterEqual(len(scripts), 1)



if __name__ == "__main__":
    unittest.main()
