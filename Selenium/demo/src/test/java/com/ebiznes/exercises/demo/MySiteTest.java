package com.ebiznes.exercises.demo;
import com.codeborne.selenide.Configuration;
import com.codeborne.selenide.Selenide;
import com.codeborne.selenide.logevents.SelenideLogger;
import io.qameta.allure.selenide.AllureSelenide;
import org.junit.jupiter.api.*;
import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.*;

import static com.codeborne.selenide.Condition.attribute;
import static com.codeborne.selenide.Condition.visible;
import static com.codeborne.selenide.Selenide.*;

public class MySiteTest {
    WebDriver driver;

    @BeforeEach
    void setup() {
        var options = new ChromeOptions();
        options.addArguments("--remote-allow-origins=*");
        driver = new ChromeDriver(options);
    }

    @Test
    public void test1() {
        driver.get("http://127.0.0.1:5173/");

        var title = driver.getTitle();

        Assertions.assertEquals(title, "Digby | Your taste matters");

        driver.quit();
    }

    @Test
    public void test2() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElement(new By.ByClassName("text-5xl")).getText();

        Assertions.assertEquals(whyAreWeDifferentText, "Why are we different?");

        driver.quit();
    }

    @Test
    public void test3() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElement(By.className("grid")).findElements(By.tagName("div")).get(0).getText();

        Assertions.assertEquals(whyAreWeDifferentText, WebsiteTexts.QUALITY_OVER_QUANTITY.getText());

        driver.quit();
    }

    @Test
    public void test4() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElement(By.className("grid")).findElements(By.tagName("div")).get(1).getText();

        Assertions.assertEquals(whyAreWeDifferentText, WebsiteTexts.ETHICALLY_SELECTED_MANUFACTURERS.getText());

        driver.quit();
    }

    @Test
    public void test5() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElement(By.className("grid")).findElements(By.tagName("div")).get(2).getText();

        Assertions.assertEquals(whyAreWeDifferentText, WebsiteTexts.YOUR_VOICE_MATTERS.getText());

        driver.quit();
    }

    @Test
    public void test6() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElement(By.className("grid")).findElements(By.tagName("div")).get(3).getText();

        Assertions.assertEquals(whyAreWeDifferentText, WebsiteTexts.WE_TAKE_CARE.getText());

        driver.quit();
    }

    //bg-[#DE972F] grid grid-cols-1 mx-auto py-10 px-3 gap-4 text-[#DE402F]
    @Test
    public void test7() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElements(By.className("bg-[#DE972F]")).get(1).getText();

        Assertions.assertEquals(whyAreWeDifferentText, WebsiteTexts.HOME_FOOTER.getText());

        driver.quit();
    }

    //        var whyAreWeDifferentText = driver.findElement(By.className("bg-[#DE972F]")).findElement(By.className("grid")).getText();
    @Test
    public void test8() {
        driver.get("http://127.0.0.1:5173/");

        var whyAreWeDifferentText = driver.findElements(By.className("bg-[#DE972F]")).get(0).getText();

        Assertions.assertEquals(whyAreWeDifferentText, "Digby\nProducts\nCart");

        driver.quit();
    }

    @Test
    public void test9() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(0).findElement(By.className("flex")).findElement(By.className("font-bold")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Zajebista koszulka");

        driver.quit();
    }

    @Test
    public void test10() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(2).findElement(By.className("font-bold")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Zajebisty zegarek");

        driver.quit();
    }

    @Test
    public void test11() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(4).findElement(By.className("font-bold")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Zaliczenie z Ebiznesu");

        driver.quit();
    }

    @Test
    public void test12() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(6).findElement(By.className("font-bold")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "3.0 z PRiRu");

        driver.quit();
    }

    //Description tests
    @Test
    public void test13() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(0).findElement(By.className("text-sm")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Nie ma co, musisz ja kupic");

        driver.quit();
    }

    @Test
    public void test14() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(2).findElement(By.className("text-sm")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Pokazuje czas lepiej niz kazdy inny");

        driver.quit();
    }

    @Test
    public void test15() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(4).findElement(By.className("text-sm")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Fajnie byloby miec co nie");

        driver.quit();
    }

    @Test
    public void test16() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(6).findElement(By.className("text-sm")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "Nieosiągalne");

        driver.quit();
    }

    //Price tests
    @Test
    public void test17() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(0).findElement(By.className("text-2xl")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "999 ZŁ");

        driver.quit();
    }

    @Test
    public void test18() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(2).findElement(By.className("text-2xl")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "69 ZŁ");

        driver.quit();
    }

    @Test
    public void test19() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(4).findElement(By.className("text-2xl")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "420 ZŁ");

        driver.quit();
    }

    @Test
    public void test20() {
        driver.get("http://127.0.0.1:5173/");

        driver.findElements(By.className("bg-[#DE972F]")).get(0).findElements(By.className("md:flex")).get(0).click();

        var zajebistaKoszulka = driver.findElements(By.className("border")).get(6).findElement(By.className("text-2xl")).getText();

        Assertions.assertEquals(zajebistaKoszulka, "0 ZŁ");

        driver.quit();
    }
}
